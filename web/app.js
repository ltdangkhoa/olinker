document.addEventListener('DOMContentLoaded', () => {
    const consoleEl = document.getElementById('console');
    
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
    document.getElementById('save-btn').addEventListener('click', () => {
        log('Configuration saved in localstorage (In production, wire to Go config API)', 'success');
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
