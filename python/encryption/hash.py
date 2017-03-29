import hashlib


def md5(s):
    m = hashlib.md5(s)
    return m.hexdigest()

appkey = '561f569367e58e036900333e'
app_master_secret = 'v31jmorwtoerqbwasmcs94blqte4nv6q'
method = 'POST'
url = 'http://msg.umeng.com/api/send'
post_body = '''{"AppKey":"561f569367e58e036900333e","TimeStamp":"1447898427","Type":"groupcast","DeviceTokens":"","AliasType":"","Alias":"","FileID":"","ProductionMode":"false","Description":"","ThirdPartyID":"","Policy":{"StartTime":"","ExpireTime":"2015-11-19 03:00:27","MaxSenNum":0,"OutBizNo":""},"Filter":{"where":{"and":[{"or":[{"tag":"cn.tsinghua"}]}]}},"Payload":{"DisplayType":"notification","Body":{"Ticker":"","Title":"a test title","Text":"blablabla...","Icon":"","LargeIcon":"","Img":"","Sound":"","PlayVibrate":"","PlayLights":"","PlaySound":"","AfterOpen":"","URL":"","Activity":"","Custom":null,"Extra":{"URL":"http://www.baidu.com","ThreadID":"thread_id","InstituteID":"institute_id","ProgramID":"program_id"}}}}'''
sign = md5('%s%s%s%s' % (method, url, post_body, app_master_secret))
print sign
