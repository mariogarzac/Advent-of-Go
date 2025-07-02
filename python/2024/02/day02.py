
def read_file():
    reports = []
    with open("test.txt") as file:
        for line in file:
            report = (line.strip()).split()
            reports.append(report)

    return reports

def part1(reports):
    count = len(reports)

    for report in reports:
        increasing = report[1] > report[0]

        for j in range(len(report) - 1):
            diff = (abs(int(report[j]) - int(report[j + 1])))
            valid_diff = 1 <= diff and diff <= 3
            is_sequential = (increasing and report[j] < report[j - 1]) or (not increasing and report[j] > report[j - 1])

            if not valid_diff or not is_sequential:
                print(report)
                count -= 1
                break

    return count

def main():
    reports = read_file()
    print(part1(reports))

main()
