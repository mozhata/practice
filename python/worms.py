import simplejson as json1
import requests
from bs4 import BeautifulSoup
import time

url = "http://www.graduate.study.cam.ac.uk/courses"
datatable = "http://gaobase.admin.cam.ac.uk/api/courses.datatable"
json = requests.get(datatable)
dicsoup = BeautifulSoup(response.text) = json1.loads(json.text)
#print len(dic['data'])


def sub_info():
    #the last 25 keywords
    keys = [u'Extra Materials', u'GRE/GMAT Score', u'Audio recordings', u'Evidence of Competence in English', u'Research Proposal', u'GRE Test (non-UK)', u'Supplementary Info Form', u'Choral Video', u'Statement of Interests', u'Field work costs', u'Personal Development Questionnaire', u'Transcript', u"Details of previous Master's", u'First Academic Reference', u'Advanced Notification Form', u'Composition Works', u'Covering Letter', u'Attainments in languages', u'Second Academic Reference', u'Personal Reference', u'CV/Resum\xe9', u'Sample of Work', u'Third Academic Reference', u'Portfolio', u'Attainments in Classics']
    #table_field
    title = ['ID', 'sub_link', 'full_name', 'department', 'full_time_length', 'degree_name', 'mail', 'course_webpage', 'term', 'year', 'deadline1name', 'deadline1', 'deadline2name', 'deadline2', 'deadline3name', 'deadline3', 'deadline4name', 'deadline4', 'deadline5name', 'deadline5', 'deadline6name', 'deadline6', 'academic_standard', 'IELTS_listening', 'IELTS_writing', 'IELTS_reading', 'IELTS_speaking', 'IELTS_total', 'toefl_listening', 'toefl_writing', 'toefl_reading', 'toefl_speaking', 'toefl_total', 'home_fee1_name',' home_fee1', 'home_fee2_name', ' home_fee2', 'home_fee_total', 'oversea_fee1_name', 'oversea_fee1', 'oversea_fee2_name', 'oversea_fee2','oversea_fee3_name', 'oversea_fee3', 'oversea_total',' application_fee',
u'Extra Materials', u'GRE/GMAT Score', u'Audio recordings', u'Evidence of Competence in English', u'Research Proposal', u'GRE Test (non-UK    )', u'Supplementary Info Form', u'Choral Video', u'Statement of Interests', u'Field work costs', u'Personal Development Questionnaire', u'Transcript', u"D    etails of previous Master's", u'First Academic Reference', u'Advanced Notification Form', u'Composition Works', u'Covering Letter', u'Attainments in langu    ages', u'Second Academic Reference', u'Personal Reference', u'CV/Resum\xe9', u'Sample of Work', u'Third Academic Reference', u'Portfolio', u'Attainments i    n Classics']

    title_str = "','".join(title)
        f = open("table.csv",'a')
        f.write("'" + title_str.encode('utf-8', 'ignore') + "'" + '\n')
        f.close()
    #some linke can not get data form them
    wired_link = []
    #there are 298 in total
    for index in range(298):
        key = []#to store 'yes'
        for i in range(25):
            key.append('')
         #two lists are created to store the repeat data
                f_space = []
                b_space = []
        f_space.append(str(index))
        #every single project's webpage
        sub_link = dic['data'][index]['prospectus_url']
        f_space.append(sub_link)
        s = sub_link.split('.')
        length = len(s[1])
        #to judge whether the linke is wired
        if length != 8:
            wired_link.append(sub_link)
            continue
        fu = dic['data'][index]['full_name']
        ful = fu.split(',')
        full_name = '/'.join(ful)
        f_space.append(full_name)
        dep = ','.join(dic['data'][index]['departments'])
        departs = dep.split(',')
        department = '/'.join(departs)
        f_space.append(department)
        print(full_name)#test the programme
        #every single overseas_fee's webpages' linke
        fee_link = "http://gaobase.admin.cam.ac.uk/api/courses/" + dic['data'][index]['code'] + "/financial_tracker.html?fee_status=O&children=0"
        response = requests.get(sub_link)
        soup = BeautifulSoup(response.text)
        fee_response = requests.get(fee_link)
        fee = BeautifulSoup(fee_response.text)
        full_time_length = soup.select('div.field-items h4 i.fa.fa-fw.fa-clock-o')[0].next_sibling.strip()
        f_space.append(full_time_length)
        degree_name = soup.select('div.field-items h4 i.fa.fa-fw.fa-graduation-cap')[0].next_sibling.next_sibling.string.strip()
        f_space.append(degree_name)
        if len(soup.select('div.field-items h4 i.fa.fa-fw.fa-envelope')) >= 1:
            mail = soup.select('div.field-items h4 i.fa.fa-fw.fa-envelope')[0].next_sibling.next_sibling.attrs['href']
        if len(soup.select('div.field-items h4 i.fa.fa-fw.fa-envelope')) < 1:
            mail = ''
        f_space.append(mail)
        course_webpage = ''
        if len(soup.select('div.field-items h4 i.fa.fa-fw.fa-link')) >= 1:
            course_webpage =  soup.select('div.field-items h4 i.fa.fa-fw.fa-link')[0].next_sibling.next_sibling.attrs['href']
        f_space.append(course_webpage)
        #if this project's application has shut down, add this to wired_link
        t = len(soup.select('div.panel.panel-default'))
        if t <= 1:
            wired_link.append(sub_link)
            continue
        term = (soup.select('div.panel.panel-default')[0].contents[1].contents[1].string.strip()).split()[0]
        year = (soup.select('div.panel.panel-default')[0].contents[1].contents[1].string.strip()).split()[1]
        deadline1name = soup.select('dl')[0].contents[1].string
        deadline1 = turn_time(soup.select('div.panel.panel-default dl')[0].contents[3].string.split())
        deadline2name = soup.select('div.panel.panel-default dl')[0].contents[5].string
        deadline2 = turn_time(soup.select('div.panel.panel-default dl')[0].select('dd')[1].string.strip().split())
        t =len(soup.select('div#funding-deadlines dl dt'))#test there are how many application are available
        fund_deadline = []
        for dt in soup.select('div#funding-deadlines dl dt'):
            #deadlinename3,4,5,6
            fund_deadline.append((dt.text).strip())
            #deadline3,4,5,6
            fund_deadline.append(turn_time((dt.next_sibling).strip().split(' ')))
        if t < 1:
            fund_deadline.append('null')
            fund_deadline.append('null' )
        if t < 2:
            fund_deadline.append('null')
            fund_deadline.append('null')
        if t < 3:
                        fund_deadline.append('null')
                        fund_deadline.append('null')
        if t < 4:
            fund_deadline.append('')
            fund_deadline.append('')

        academic_standard = soup.select('div#requirements hr')[0].previous_element.previous_element.previous_element.string
        b_space.append(academic_standard)
        #to check whether threr is a table for toefl
        length =len(soup.select('div.campl-column6 table.campl-table.campl-table-condensed.campl-table-bordered'))
        IELTS_listening =  soup.select('div.campl-column6 table.campl-table.campl-table-condensed.campl-table-bordered')[0].select('td')[0].string.strip()
        b_space.append(IELTS_listening)
        IELTS_reading =  soup.select('div.campl-column6 table.campl-table.campl-table-condensed.campl-table-bordered')[0].select('td')[2].string.strip()
        b_space.append(IELTS_reading)
        IELTS_speaking =  soup.select('div.campl-column6 table.campl-table.campl-table-condensed.campl-table-bordered')[0].select('td')[3].string.strip()
        b_space.append(IELTS_speaking)
        IELTS_writing =  soup.select('div.campl-column6 table.campl-table.campl-table-condensed.campl-table-bordered')[0].select('td')[1].string.strip()
        b_space.append(IELTS_writing)
        IELTS_total =  soup.select('div.campl-column6 table.campl-table.campl-table-condensed.campl-table-bordered')[0].select('td')[4].string.strip()
        b_space.append(IELTS_total)
        toefl_listening = ''
        toefl_writing = ''
        toefl_reading = ''
        toefl_speaking = ''
        toefl_total = ''
        if length > 1:
            toefl_listening = soup.select('div.campl-column6 table.campl-table.campl-table-condensed.campl-table-bordered')[1].contents[3].contents[1].contents[3].string
#           b_space.append(toefl_listening)
            toefl_writing = soup.select('div.campl-column6 table.campl-table.campl-table-condensed.campl-table-bordered')[1].contents[3].contents[3].contents[3].string
#           b_space.append(toefl_writing)
            toefl_reading = soup.select('div.campl-column6 table.campl-table.campl-table-condensed.campl-table-bordered')[1].contents[3].contents[5].contents[3].string
#           b_space.append(toefl_reading)
            toefl_speaking = soup.select('div.campl-column6 table.campl-table.campl-table-condensed.campl-table-bordered')[1].contents[3].contents[7].contents[3].string
#           b_space.append(toefl_speaking)
            toefl_total = soup.select('div.campl-column6 table.campl-table.campl-table-condensed.campl-table-bordered')[1].contents[3].contents[9].contents[3].string
#           b_space.append(toefl_total)
        b_space.append(toefl_listening)
        b_space.append(toefl_writing)
        b_space.append(toefl_reading)
        b_space.append(toefl_speaking)
        b_space.append(toefl_total)
        #if there is no oversea_fee,programme will fail ,and I can't fix it ,so append it into wired_link[]
        try:
            length1 = len(soup.select('div#fee_1 table')[0].select('tbody tr'))
            length0 = 0
            if len(fee.select('div#fee_1 table')) > 0:
                length0 = len(fee.select('div#fee_1 table')[0].select('tbody tr'))
            if length0 >= 1:
                length2 = len(fee.select('div#fee_1 table')[0].select('tbody tr'))
            home_fee1_name = soup.select('div#fee_1 table')[0].contents[3].contents[1].contents[1].string
            b_space.append(home_fee1_name)
            home_fee1 = ''.join(soup.select('div#fee_1 table')[0].contents[3].contents[1].contents[3].string.split(','))[1:].strip()
            b_space.append(home_fee1)
            if length1 >= 2:
                home_fee2_name = soup.select('div#fee_1 table')[0].contents[3].contents[3].contents[1].string
                home_fee2 = ''.join(soup.select('div#fee_1 table')[0].contents[3].contents[1].contents[3].string.split(','))[1:].strip()
            if length1 < 2:
                home_fee2_name = ''
                home_fee2 = ''
            b_space.append(home_fee2_name)
            b_space.append(home_fee2)
            home_fee_total = ''.join(soup.select('div#fee_1 table')[0].contents[5].contents[1].contents[3].string.split(','))[1:].strip()
            b_space.append(home_fee_total)
            if length2 >= 1:
                oversea_fee1_name = fee.select('table')[0].contents[3].contents[1].contents[1].string
                oversea_fee1 = ''.join(fee.select('table')[0].contents[3].contents[1].contents[3].string.split(','))[1:].strip()
            if length2 < 1:
                oversea_fee1_name = ''
                oversea_fee1 = ''
            b_space.append(oversea_fee1_name)
            b_space.append(oversea_fee1)
            if length2 >= 2:
                oversea_fee2_name = fee.select('table')[0].contents[3].contents[3].contents[1].string
                        oversea_fee2 = ''.join(fee.select('table')[0].contents[3].contents[3].contents[3].string.split(','))[1:].strip()
            if length2 < 2:
                oversea_fee2_name = ''
                oversea_fee2 = ''
            b_space.append(oversea_fee2_name)
            b_space.append(oversea_fee2)
            if length2 >= 3:
                oversea_fee3_name = fee.select('table')[0].contents[3].contents[5].contents[1].string
                oversea_fee3 = ''.join(fee.select('table')[0].contents[3].contents[5].contents[3].string.split(','))[1:].strip()
            if length2 < 3:
                oversea_fee3_name = ''
                oversea_fee3 = ''
            b_space.append(oversea_fee3_name)
            b_space.append(oversea_fee3)
            oversea_total = ''.join(fee.select('table')[0].contents[5].contents[1].contents[3].string.split(','))[1:].strip()
            b_space.append(oversea_total)
        except:
            wired_link.append(sub_link)
            continue
        #tem[] is about to collect keywords in this page
                t0 = len(soup.select('div#new_checklist table')[0].select('tbody tr strong'))
                tem = []
                for i in range(t0):
                        tem.append(soup.select('div#new_checklist table')[0].select('tbody tr strong')[i].string.strip())
        #keys[] are the list to store all the keywords in all,
        #key[] are empty ,to store 'yes'
                for i in range(len(keys)):
                        for j in range(len(tem)):
                                if keys[i] == tem[j]:
                                        key[i] = 'yes'
                                        break
                                else:
                                        continue
        application_fee = ''.join(soup.select('div#apply strong')[1].string[1:-15].strip().split(','))
        b_space.append(application_fee)
        b_space.extend(key)
        #fields is to store all this stuff in one single loop
        fields = [str(index), sub_link, full_name, department, full_time_length, degree_name, mail, course_webpage, term, year, deadline1name, deadline1, deadline2name, deadline2, fund_deadline[0], fund_deadline[1], fund_deadline[2], fund_deadline[3], fund_deadline[4], fund_deadline[5], fund_deadline[6], fund_deadline[7], academic_standard, IELTS_listening, IELTS_writing, IELTS_reading, IELTS_speaking, IELTS_total, toefl_listening, toefl_writing, toefl_reading, toefl_speaking, toefl_total, home_fee1_name, home_fee1, home_fee2_name, home_fee2, home_fee_total, oversea_fee1_name, oversea_fee1, oversea_fee2_name, oversea_fee2,oversea_fee3_name, oversea_fee3,oversea_total, application_fee]
        fields.extend(key)
        fields.append('\n')
#       print(f_space[:3])
#       print(fields[-3:])
##      print(len(fields))
        # to check if there is more than one 'term' , if the answer is yes ,do something
        t = len(soup.select('div.panel.panel-default'))
        if (t-1) < 2:
            lines = fields
        if (t-1) >= 2:
            e = t-1
            for de in range(1,e):
                #the second line which is different
                mass = []
                mass.append(soup.select('div.panel.panel-default')[de].select('h4')[0].string.strip().split()[0])
                mass.append(soup.select('div.panel.panel-default')[de].select('h4')[0].string.strip().split()[1])
                mass.append(soup.select('div.panel.panel-default')[de].select('dt')[0].string.strip())
                print(t, e, de)#for test
                timelis = soup.select('div.panel.panel-default')[1].select('dd')[0].string.strip().split()
                print(timelis)
                tim = turn_time(timelis)
                print(tim)
                mass.append(tim)
                mass.append(soup.select('div.panel.panel-default')[de].select('dt')[1].string.strip())
                timelis = soup.select('div.panel.panel-default')[1].select('dd')[0].string.strip().split()
                tim = turn_time(timelis)
                mass.append(tim)
                mass.extend(fund_deadline)

                #build the second line
                mass.extend(b_space)
                f_space.extend(mass)
                f_space.append('\n')
            fields.extend(f_space)
            del fields[72]
#           print(fields[70:74])

        field_str = ",".join(fields)
        f = open("table.csv",'a')
        f.write(field_str.encode('utf-8', 'ignore'))
        f.close()
        #to test if the loop is done
        print(full_name)
        # aimed to fix the oversea_fee's problem ,but failed
        length0 = 0
    #get the wired link
    print(wired_link)
# for example ,turn "['Sep','1,','2014']" to '2014-09-01'
def turn_time(daytime):
    b = daytime[0][:3]#month
    j = daytime[1][:-1]#day
    if len(j.strip())==1:
        j = '0' + j
    Y = daytime[2]#year
    t = Y + '-' + j + '-' + b
    mkt = time.mktime(time.strptime(t,'%Y-%d-%b'))
    x = time.localtime(mkt)
    time_format =  time.strftime('%Y-%m-%d',x)
    return time_format

sub_info()
print('done')
