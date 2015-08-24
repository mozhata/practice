# move file to folder with the same name
import os
import shutil
# for filename in os.listdir(r'c:\windows'):
#     print filename
paths = "/home/zyk/test"

# for f in os.listdir(path):
# 	print("test")
# 	if os.path.isfile(f):
# 		print(f)
files = os.listdir(paths)
print(files)
# for fi in files:
# 	print(os.path.isdir(os.path.join(paths, fi)), os.path.isfile(os.path.join(paths, fi)))
	# print(os.path.isdir(fi), os.path.isfile(fi))

folders = []
for f in files:
	if os.path.isdir(os.path.join(paths, f)):
		folders.append(f)

print(folders)
filess = []
for f in files:
	if os.path.isfile(os.path.join(paths, f)):
		filess.append(f)

print(filess)

# for f in filess:
# 	print(f.split(".")[0])

# print("a" in folders)
# for f in filess:
# 	if f.split(".")[0] in folders:
# 		folder = os.path.join(paths, f.split(".")[0])
# 		current_file = os.path.join(paths, f)
# 		dist_file = folder + current_file
# 		print(folder, current_file, os.path.isdir(folder), os.path.isfile(current_file))
# 		shutil.copy(current_file, dist_file)
# 		os.remove(current_file)
# print(os.listdir(paths))


def copyfile(dist, f):
	fd = open(dist)
	fs = open(f)
	fd.write(fs.read())
	fs.close()
	fd.close()


for f in filess:
	dist = paths + f.split(".")[0] + "/" + f
	# print(dist)
	copyfile(dist, f)
