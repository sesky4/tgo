import sys
import os

def git_tag(tag):
	os.system("git tag %s" % (dir, tag))

def main():
	tag = sys.argv[1]

	git_tag(tag)

	for svc in os.listdir("tencentcloud"):
		svc_tag = 'tencentcloud/%s/%s' % (svc, tag)
		git_tag(base_dir, svc_tag), 

if __name__ == "__main__":
	main()
