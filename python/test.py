# coding=utf-8
# birth = input('birth')
# if int(birth) > 2000:
# 	print('00前')
# else:
# 	print("00后")
# print(int("2"))
# # print(int("a"))
# if int("a"):
# 	print("yes")
l = [
	"architecture",
	"area_studies",
	"biological_agricultural_engineering",
	"biology",
	"biomedical_engineering",
	"chemical_engineering",
	"chemistry",
	"civil_engineering",
	"communication",
	"computer_engineering",
	"computer_science",
	"criminology",
	"economics",
	"education",
	"electrical_engineering",
	"english",
	"entrepreneurship",
	"environmental_engineering",
	"environmental_science",
	"epidemiology",
	"finance",
	"financial_engineering",
	"fine_arts",
	"geosciences",
	"health",
	"history",
	"human_resource_management",
	"industrial_engineering",
	"information_systems",
	"international_business",
	"language",
	"law",
	"literature",
	"management",
	"marketing",
	"materials",
	"mathematics",
	"mba",
	"mechanical_engineering",
	"medicine",
	"nuclear_engineering",
	"operations",
	"pharmacy",
	"philosophy",
	"physics",
	"political_science",
	"psychology",
	"public_affairs",
	"public_health",
	"public_management_administration",
	"public_policy_analysis",
	"religion",
	"social_work",
	"sociology",
	"statistics",
	"strategy",
	"supply_chain_and_logistics",
	"urban_planning",
]
print(len(l))
for fos in l:
	print('PERMISSION_%s_USER = \'%s_user\'' % (fos.upper(), fos))
print("~~~~~")
for fos in l:
	print('PERMISSION_%s_USER,' % fos.upper())
print("~~~~~~~")
for fos in l:
	print('''
@property
def is_fos_%s_user(self):
	return self.has_permission(self.PERMISSION_%s_USER)''' % (fos, fos.upper()))
