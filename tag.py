import sys
import os

def git_tag(tag):
	os.system("git tag %s" % tag)

def main():
	tag = sys.argv[1]
	print tag[0]

	git_tag(tag)

	for svc in os.listdir("tencentcloud"):
		git_tag('tencentcloud/%s/%s' % (svc, tag))

if __name__ == "__main__":
	main()
