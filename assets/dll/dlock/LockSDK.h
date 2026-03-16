#ifndef __LOCK_SDK_H__
#define __LOCK_SDK_H__
   


// 卡片错误 
enum ERROR_TYPE
{
    OPR_OK              =   1,      // 操作成功
    NO_CARD			    =   -1,     // 没检测到卡片
    NO_RW_MACHINE	    =   -2,     // 没检测到读卡器
    INVALID_CARD	    =   -3,     // 无效卡
    CARD_TYPE_ERROR	    =   -4,     // 卡类型错误
    RDWR_ERROR		    =   -5,     // 读写错误
    PORT_NOT_OPEN	    =   -6,     // 端口未打开
    END_OF_DATA_CARD    =   -7,     // 数据卡结束
    INVALID_PARAMETER   =   -8,     // 无效参数
    INVALID_OPR		    =   -9,     // 无效操作
    OTHER_ERROR		    =   -10,    // 其它错误
    PORT_IN_USED	    =   -11,    // 端口已被占用
    COMM_ERROR		    =   -12,    // 通讯错误    
    ERR_CLIENT          =   -20,    // 客户码错误    
    ERR_NOT_REGISTERED  =   -29,    // 未注册
    ERR_NO_CLIENT_DATA  =   -30,     // 无授权卡信息
    ERR_ROOMS_CNT_OVER  =   -31,    // 房数超出了可用扇区
}; 


#ifdef __cplusplus
	extern "C" { 
#endif

/*=============================================================================
函数名：                        TP_Configuration
;
功　能：动态库初始化配置, 完成门锁类型选择/发卡器连接等
输  入：lock_type -- 门锁类型(也就是使用的卡片类型): 4-RF57门锁; 5-RF50门锁
输  出: 无
返回值：错误类型
=============================================================================*/
int __stdcall TP_Configuration(int lock_type);


/*=============================================================================
函数名：                        TP_MakeGuestCard
;
功　能：制作宾客卡
输  入：room_no         --  房号:       字符串, 例如 "001.002.00003.A"
        checkin_time    --  入住时间：  年月日时分秒, 字符串格式 "YYYY-MM-DD hh:mm:ss"
        checkout_time   --  预离时间：  年月日时分秒, 字符串格式 "YYYY-MM-DD hh:mm:ss"
        iflags          --  宾客卡选项, 参见Defines中的GUEST_FLAGS定义,一般置0
输  出: card_snr        -- 卡号:        字符串, 至少预分配20字节
例  子: Room="001.002.00003.A", SDateTime="2008-06-06 12:30:59", EDateTime="2008-06-07 12:00:00"
        iFlags=0
返回值：错误类型
=============================================================================*/
int __stdcall TP_MakeGuestCard(char *card_snr, char *room_no, char *checkin_time,char *checkout_time, int iflags);



/*=============================================================================
函数名：                        TP_ReadGuestCard
;
功　能：读宾客卡信息
输  入：无。
输  出: card_snr        --  卡号:       字符串, 至少预分配20字节
        room_no         --  房号:       字符串, 至少预分配20字节
        checkin_time    --  入住时间：  年月日时分秒, 字符串格式 "YYYY-MM-DD hh:mm:ss", 至少预分配30字节
        checkout_time   --  预离时间：  年月日时分秒, 字符串格式 "YYYY-MM-DD hh:mm:ss", 至少预分配30字节
返回值：错误类型
=============================================================================*/
int __stdcall	TP_ReadGuestCard(char *card_snr,char *room_no, char *checkin_time, char *checkout_time);


/*=============================================================================
函数名：                        TP_CancelCard
;
功　能：注销卡片/卡片回收
输  入: 无
输  出：
输  出: card_snr    -- 卡号: 字符串, 至少预分配20字节
返回值：错误类型
=============================================================================*/
int __stdcall TP_CancelCard(char *card_snr);

/*=============================================================================
函数名：                        TP_GetCardSnr
;
功　能：读取卡号(卡片的唯一的序列号)
输  入: 无
输  出: card_snr    --  卡号: 字符串, 至少预分配20字节
返回值：错误类型
=============================================================================*/
int __stdcall TP_GetCardSnr(char *card_snr);


#ifdef __cplusplus
   }
#endif

#endif              // __LOCK_SDK_H__