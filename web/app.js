document.addEventListener('DOMContentLoaded', async () => {
    const consoleEl = document.getElementById('console');
    const vendorSelect = document.getElementById('vendor-select');
    const settingsContainer = document.getElementById('vendor-settings');
    let currentConfig = {};

    const vendorFields = {
        prousb: [
            { id: 'port', label: 'USB Port', type: 'select', options: ['USB', 'proUSB'] },
            { id: 'hotel_id', label: 'Hotel ID', type: 'text' },
            { id: 'card_no', label: 'Current Card No (Auto-increment)', type: 'text', default: '0' },
            { id: 'mapping', label: 'Mapping (room_name=LockNo)', type: 'textarea' }
        ],
        dlock: [
            { id: 'lock_type', label: 'Card Type', type: 'select', options: ['RF50', 'RF57'] },
            { id: 'mapping', label: 'Mapping (RoomName=LockNumber)', type: 'textarea' }
        ],
        adel: [
            { id: 'db_server', label: 'DB Server', type: 'text', default: '\\\\SQLEXPRESS' },
            { id: 'model', label: 'Model', type: 'select', options: ['A90', 'A92', 'Lock9200', 'Lock9200T'] },
            { id: 'port', label: 'Port', type: 'select', options: ['USB', 'COM1', 'COM2', 'COM3', 'COM4'] },
            { id: 'encoder', label: 'Encoder', type: 'select', options: ['Manual Encoder', 'Automatic Encoder', 'MSR206'] },
            { id: 'encoder_tm', label: 'Encoder TM', type: 'select', options: ['DS9097E', 'DS9097U'] },
            { id: 'mapping', label: 'Mapping (room_name=room_no)', type: 'textarea' }
        ],
        hune: [
            { id: 'com', label: 'COM Port', type: 'text' },
            { id: 'nblock', label: 'NBlock', type: 'text' },
            { id: 'encrypt', label: 'Encrypt', type: 'text' },
            { id: 'card_pass', label: 'Card Pass', type: 'text' },
            { id: 'system_code', label: 'System Code', type: 'text' },
            { id: 'hotel_code', label: 'Hotel Code', type: 'text' },
            { id: 'mapping', label: 'Mapping (room_name=Address)', type: 'textarea' }
        ],
        orbita: [{ id: 'mapping', label: 'Mapping', type: 'textarea' }],
        betech: [{ id: 'mapping', label: 'Mapping', type: 'textarea' }],
        mock: [{ id: 'foo', label: 'Mock Setting', type: 'text' }]
    };

    function renderFields(vendor) {
        settingsContainer.innerHTML = '';
        const fields = vendorFields[vendor] || [];
        fields.forEach(f => {
            const div = document.createElement('div');
            div.className = 'mb-3';
            const val = currentConfig.settings ? (currentConfig.settings[f.id] || f.default || '') : (f.default || '');
            
            let inputHtml = '';
            if (f.type === 'select') {
                inputHtml = `<select class="form-select custom-select" data-id="${f.id}">
                    ${f.options.map(o => `<option value="${o}" ${o === val ? 'selected' : ''}>${o}</option>`).join('')}
                </select>`;
            } else if (f.type === 'textarea') {
                inputHtml = `<textarea class="form-control custom-input" data-id="${f.id}" rows="3">${val}</textarea>`;
            } else {
                inputHtml = `<input type="text" class="form-control custom-input" data-id="${f.id}" value="${val}">`;
            }

            div.innerHTML = `<label class="form-label">${f.label}</label>${inputHtml}`;
            settingsContainer.appendChild(div);
        });
    }

    async function loadConfig() {
        try {
            const res = await fetch('/config');
            currentConfig = await res.json();
            document.getElementById('api-port').value = currentConfig.port;
            document.getElementById('dll-path').value = currentConfig.dll_path;
            vendorSelect.value = currentConfig.vendor;
            renderFields(currentConfig.vendor);
            log('Config loaded from server', 'system');
        } catch (e) {
            log('Failed to load config: ' + e.message, 'error');
        }
    }

    vendorSelect.addEventListener('change', (e) => {
        renderFields(e.target.value);
    });

    await loadConfig();

    function log(message, type = 'system') {
        const time = new Date().toLocaleTimeString();
        const entry = document.createElement('div');
        entry.className = `log-entry ${type}`;
        entry.innerHTML = `<span style="opacity:0.5">[${time}]</span> ${message}`;
        consoleEl.insertBefore(entry, consoleEl.firstChild);
    }

    // Tabs logic
    document.querySelectorAll('.tab').forEach(tab => {
        tab.addEventListener('click', (e) => {
            document.querySelectorAll('.tab, .tab-content').forEach(el => el.classList.remove('active'));
            tab.classList.add('active');
            document.getElementById(tab.dataset.target).classList.add('active');
        });
    });

    // Form settings logic
    document.getElementById('save-btn').addEventListener('click', async () => {
        const settings = {};
        settingsContainer.querySelectorAll('[data-id]').forEach(el => {
            settings[el.dataset.id] = el.value;
        });

        const newConfig = {
            vendor: vendorSelect.value,
            dll_path: document.getElementById('dll-path').value,
            port: parseInt(document.getElementById('api-port').value),
            settings: settings
        };

        try {
            const res = await fetch('/config', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(newConfig)
            });
            if (res.ok) {
                log('Configuration saved successfully!', 'success');
                currentConfig = newConfig;
            } else {
                throw new Error(await res.text());
            }
        } catch (e) {
            log('Failed to save config: ' + e.message, 'error');
        }
    });

    // Encode Test
    const testEncodeBtn = document.getElementById('test-encode-btn');
    testEncodeBtn.addEventListener('click', async () => {
        const vendor = document.getElementById('vendor-select').value;
        const room = document.getElementById('test-room').value || '101';
        log(`Sending encode request for room ${room} to /${vendor}/write_card...`, 'system');
        
        try {
            const res = await fetch(`/${vendor}/write_card`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ 
                    room_name: room, 
                    BeginTime: new Date().toISOString(), 
                    EndTime: new Date(Date.now() + 86400000).toISOString() 
                })
            });
            const data = await res.json();
            
            if (res.ok) {
                log(`Success: Encoded Card [${data.card_no}] for Room [${data.room_name}]`, 'success');
            } else {
                throw new Error(data.message || 'Server error');
            }
        } catch (err) {
            log(`Error: ${err.message}`, 'error');
        }
    });

    // Read test
    const testReadBtn = document.getElementById('test-read-btn');
    testReadBtn.addEventListener('click', async () => {
        const vendor = document.getElementById('vendor-select').value;
        log(`Reading card from /${vendor}/read_card...`, 'system');
        try {
            const res = await fetch(`/${vendor}/read_card`, { method: 'POST' });
            const data = await res.json();
            log(`Card Read: ${JSON.stringify(data)}`, 'success');
        } catch (err) {
            log(`Error: ${err.message}`, 'error');
        }
    });
});
