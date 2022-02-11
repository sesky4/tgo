import sys
import os
import glob

def main():
    ver = sys.argv[1]
    for p in glob.glob('tencentcloud/*/go.mod'):
        if p == 'tencentcloud/common/go.mod':
            continue

        with open(p, 'r+') as f:
            c = ''
            replaced = False
            for line in f.readlines():
                if 'require github.com/sesky4/tgo/tencentcloud/common' not in line:
                    c += line
                    continue

                parts = [x.strip() for x in line.split(" ")]
                c += "%s %s %s\n" % (parts[0], parts[1], ver)
                replaced = True

            if not replaced:
                c += "\nrequire github.com/sesky4/tgo/tencentcloud/common %s\n" % ver

            f.seek(0)
            f.truncate(0)
            f.write(c)

    with open('go.mod', 'r+') as f:
        c = ''
        for line in f.readlines():
            if 'github.com/sesky4/tgo/tencentcloud/' not in line:
                c += line
                continue

            parts = [x.strip() for x in line.split(" ")]
            c += "\t%s %s\n" % (parts[0], ver)

        f.seek(0)
        f.truncate(0)
        f.write(c)

if __name__ == "__main__":
	main()
