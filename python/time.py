def format_time(time):
    ts = time.split('T')
    ts1 = ts[0].split('-')
    ts2 = ts[1].split(':')
    for i in range(1, len(ts1)):
        if len(ts1[i]) == 1 and int(ts1[i]) < 10:
            ts1[i] = "0" + ts1[i]
    date = '-'.join(ts1)
    for i in range(len(ts2)):
        if len(ts2[i]) == 1 and int(ts2[i]) < 10:
            ts2[i] = "0" + ts2[i]
    time = ':'.join(ts2)
    return 'T'.join([date, time])

t = '2015-10-09T07:21:09'
print(len(t))

# t2 = '2015-10-9T7:21:9'

# ts = format_time(t)
# print("ts: ", ts)
# ts2 = format_time(t2)
# print("ts2: ", ts2,"\nt2: ", t2)
dic = {"a" : "apple", "b" : "banana", "g" : "grape", "o" : "orange"}
print dic.items()