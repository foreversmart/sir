// 新建组件辅助

const fs = require("fs")
const path = require("path")

const name = process.argv[2]
if (!name) {
  console.log("给个组件名字")
  return
}
const dir = process.argv[3]
if (!dir) {
  console.log("给个组件目录")
  return
}

const cname = name[0].toUpperCase() + name.substr(1)

const fileComponent = `package command

import (
	cli "gopkg.in/urfave/cli.v1"
)

var Cmd${cname} = cli.Command{
  Name: "${name}",
  UsageText: "",
	Category:  "",
	Usage:     "",
	Action: Action${cname},
}

func Action${cname}(c *cli.Context) error {
	return nil
}
`

const thePath = path.join(__dirname, `${dir}`)
console.log(thePath)

fs.writeFile(path.join(thePath, `${name}.go`), fileComponent, err => {
  if (err) {
    throw err
  }
  console.log("component created")
})

console.log("done.")
