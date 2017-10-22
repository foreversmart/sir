const elTaskList = document.getElementById("task-list")
const elTaskTitle = document.getElementById("task-title")
const elTaskActions = document.getElementById("task-actions")
const elTaskBody = document.getElementById("task-body")

let tasks = []
let selectedIndex = 0

bind()
loadTasks()

function loadTasks() {
  fetch("/task")
    .then(res => {
      return res.json()
    })
    .then(ret => {
      return ret.data
    })
    .then(data => {
      tasks = data
      renderTasks()
      renderDetail()
    })
}

function renderTasks() {
  const html = []
  for (let i = 0; i < tasks.length; i++) {
    html.push(`
<li data-index=${i} class="${i == selectedIndex ? "active" : ""}">
  <input type="checkbox" ${i == selectedIndex ? "checked" : ""} />
  <div class="preview">
    <h3>${tasks[i].name}
      <small>${tasks[i].c_time}</small>
    </h3>
    <p>
      ${tasks[i].cmd}</p>
  </div>
</li>
  `)
  }

  elTaskList.innerHTML = html.join("")
}

function renderDetail() {
  const t = tasks[selectedIndex]
  elTaskTitle.innerHTML =
    tasks[selectedIndex].name + " [" + tasks[selectedIndex].cmd + "]"

  elTaskActions.innerHTML = `
<p>
  <a href="" class="user">Start</a>
  <a href="" class="user">Stop</a>
  <a href="" class="user">Restart</a>
  <span class="date">Created Time: ${t.c_time}</span>
</p> 
  `

  elTaskBody.innerHTML = `
  <table class="table">
    <tr>
      <td>NAME</td>
      <td>${t.name}</td>
    </tr>
  <tr>
    <td>PID</td>
    <td>${t.pid}</td>
  </tr>
  <tr>
    <td>CPU</td>
    <td>${t.cpu}</td>
  </tr>
  <tr>
    <td>MEMORY</td>
    <td>${t.mem}</td>
  </tr>
  <tr>
    <td>LOAD</td>
    <td>${t.load}</td>
  </tr>
  <tr>
    <td>STAT</td>
    <td>${t.stat}</td>
  </tr>
  <tr>
    <td>UP TIME</td>
    <td>${t.up_time}</td>
  </tr>
    <tr>
      <td>EXEC</td>
      <td>${t.cmd}</td>
    </tr>
    <tr>
      <td>WATCH</td>
      <td>${t.watch}</td>
    </tr>
    <tr>
      <td>WATCH DIR</td>
      <td>${t.watch_dir}</td>
    </tr>
    <tr>
      <td>USER</td>
      <td>${t.user}</td>
    </tr>
    <tr>
      <td>GROUP</td>
      <td>${t.group}</td>
    </tr>
    <tr>
      <td>PRIORITY</td>
      <td>${t.priority}</td>
    </tr>
    <tr>
      <td>AUTO START</td>
      <td>${t.auto_start}</td>
    </tr>
    <tr>
      <td>ERR LOG</td>
      <td>${t.err_log}</td>
    </tr>
    <tr>
      <td>STD LOG</td>
      <td>${t.std_log}</td>
    </tr>
  </table> 
  `
}

function selectIndex(index) {
  selectedIndex = index
  renderTasks()
  renderDetail()
}

function bind() {
  eventDelegate("#task-list", "li", "click", function() {
    const index = this.dataset["index"]
    selectIndex(index)
  })
}

function eventDelegate(parentSelector, targetSelector, events, foo) {
  // 触发执行的函数
  function triFunction(e) {
    // 兼容性处理
    var event = e || window.event

    // 获取到目标阶段指向的元素
    var target = event.target || event.srcElement

    // 获取到代理事件的函数
    var currentTarget = event.currentTarget

    // 处理 matches 的兼容性
    if (!Element.prototype.matches) {
      Element.prototype.matches =
        Element.prototype.matchesSelector ||
        Element.prototype.mozMatchesSelector ||
        Element.prototype.msMatchesSelector ||
        Element.prototype.oMatchesSelector ||
        Element.prototype.webkitMatchesSelector ||
        function(s) {
          var matches = (this.document || this.ownerDocument).querySelectorAll(
              s
            ),
            i = matches.length
          while (--i >= 0 && matches.item(i) !== this) {}
          return i > -1
        }
    }

    // 遍历外层并且匹配
    while (target !== currentTarget) {
      // 判断是否匹配到我们所需要的元素上
      if (target && target.matches && target.matches(targetSelector)) {
        var sTarget = target
        // 执行绑定的函数，注意 this
        foo.call(sTarget, Array.prototype.slice.call(arguments))
      }

      target = target.parentNode
    }
  }

  // 如果有多个事件的话需要全部一一绑定事件
  events.split(".").forEach(function(evt) {
    // 多个父层元素的话也需要一一绑定
    Array.prototype.slice
      .call(document.querySelectorAll(parentSelector))
      .forEach(function($p) {
        $p.addEventListener(evt, triFunction)
      })
  })
}
