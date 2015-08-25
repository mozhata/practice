# move file to folder with the same name
import os
import shutil

paths = "/home/kang/test"

files = os.listdir(paths)
print("all files: ", files)

folders = []
for f in files:
	if os.path.isdir(os.path.join(paths, f)):
		folders.append(f)

print("folders: ", folders)
filess = []
for f in files:
	if os.path.isfile(os.path.join(paths, f)):
		filess.append(f)

# print(filess)

# os.makedirs('d:/assist/set')
# os.path.exists('d:/assist/set')
# shutil.move("myfile1.txt", "../")

def move_files():
	for f in filess:
		folder = os.path.join(paths, f.split(".")[0])
		print(folder)
		if not os.path.exists(folder):
			os.makedirs(folder)
		shutil.move(os.path.join(paths, f), folder)



# def copyfile(dist, f):
# 	fd = open(dist)
# 	fs = open(f)
# 	fd.write(fs.read())
# 	fs.close()
# 	fd.close()


# def mov():
# 	for f in filess:
# 		dist = paths + f.split(".")[0] + "/" + f
# 		# print(dist)
# 		copyfile(dist, f)

move_files()