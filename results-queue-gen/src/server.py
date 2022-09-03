from collections import namedtuple

filepath = '/data/IPIP300.dat'

ParsedRow = namedtuple('ParsedRow', 'case_id sex age country answers')
positions = {
    'case_id': [0, 6],
    'sex': [6, 7],
    'age': [7, 9],
    'country': [22, 32],
    'answers': [33, 333]}

ps_records = []

with open(filepath) as fp:
    for cnt, line in enumerate(fp):
       parsed = ParsedRow(
           case_id=line[0:6].strip(),
           sex=line[6:7].strip(),
           age=line[7:9].strip(),
           country=line[22:32].strip(),
           answers=line[33:333].strip()
       )
       ps_records.append(parsed)

print(len(ps_records))
# for ps in ps_records:
#     print(len(ps_records))
