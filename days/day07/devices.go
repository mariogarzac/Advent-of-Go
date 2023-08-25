package seven

import (
	"Advent/utils"
	"fmt"
	"strconv"
	"strings"
)

func Solve(filename string){

    root := MapFileSystem(filename)
    sumDirectories(root)

    // Part 1
    fmt.Println(findSmallestDirs(root))

    // Part 2
    unusedSpace := 70000000 - root.Value
    fmt.Println(makeSpace(root, unusedSpace, 70000000))
}

func makeSpace(root *utils.TreeNode, unusedSpace int, smallest int) int {

    if unusedSpace + root.Value >= 30000000 && root.Value < smallest {
        smallest = root.Value
    }

    for _, child := range root.Children {
        smallest = makeSpace(child, unusedSpace, smallest)
    }

    return smallest
}

func findSmallestDirs(root *utils.TreeNode) int {
    sum := 0
	if root != nil {
        if root.Parent != nil && !root.IsFile {
            if root.Value < 100000 {
                sum += root.Value
            }
        }
		for _, child := range root.Children {
			sum += findSmallestDirs(child)
		}
	}
    return sum
}

func sumDirectories(root *utils.TreeNode) {

    sum := 0
    for _, child := range root.Children {
        sum += child.Value
        sumDirectories(child)
        child.Parent.Value += child.Value
    }
}

func MapFileSystem(filename string) *utils.TreeNode {

    sc, file, _ := utils.OpenFile(filename)

    defer file.Close()

    var item string
    var currParent, newChild *utils.TreeNode

    // Jump a line and create root
    sc.Scan()
    root := utils.NewTree("root", false, 0)
    currParent = root.Root

    for sc.Scan(){
        cmd := strings.Fields(sc.Text())

        // Remove the '$' from the command 
        if item = cmd[0]; item == "$" { item = cmd[1] }

        switch item {
        case "ls":
            // Ignore
        case "cd":
            // Check if the command is cd .. or cd <dir>
            // if it is .. set the current to the child's parent
            // if it is to cd, find the dir that matches the name
            if strings.Contains(sc.Text(), ".."){
                if currParent = currParent.Parent; currParent == nil {
                    currParent = root.Root
                }
            }else{
                for i := 0; i < len(currParent.Children); i++ {
                   if currParent.Children[i].Name == cmd[2] {
                       currParent = currParent.Children[i]
                   }
                }
            }
        case "dir":
            dirName := string(cmd[1])
            newChild = utils.NewTreeNode(dirName, false, 0)
            currParent.AddChild(newChild)

        default:
            // variables for readability
            fileSize, _ := strconv.Atoi(cmd[0])
            fileName := string(cmd[1])

            // create new file node
            file := utils.NewTreeNode(fileName, true, fileSize)
            currParent.AddChild(file)
        }
    }
    return root.Root
}
