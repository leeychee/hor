/**
 * Created by Macintosh on 6/21/17.
 */
let Rect = {
    //当前正在画的矩形对象
    obj: null,
    //画布
    container: null,
    //初始化函数
    init: function (containerId) {
        Rect.container = document.getElementById(containerId);
        if (Rect.container) {
            //鼠标按下时开始画
            Rect.container.onmousedown = Rect.start;
        }
        else {
            alert('You should specify a valid container!');
        }
    },
    start: function (e) {
        let o = Rect.obj = document.createElement('div');
        o.style.position = "absolute";
        // mouseBeginX，mouseBeginY是辅助变量，记录下鼠标按下时的位置
        o.className = "resize-drag";
        o.style.color = "red";
        o.mouseBeginX = Rect.getEvent(e).x;
        o.style.left = Rect.getEvent(e).layerX + 15 + 'px';
        o.mouseBeginY = Rect.getEvent(e).y;
        o.style.top = Rect.getEvent(e).clientY + 'px';
        o.style.height = 0;
        o.style.width = 0;
        o.style.border = "dotted yellow 1px";

        o.style.boxSizing="border-box";
        //向o添加一个叉叉，点击叉叉可以删除这个矩形
        let deleteLink = document.createElement('a');
        deleteLink.href = "#";
        deleteLink.onclick = function () {
            Rect.container.removeChild(this.parentNode);
            //this.parentNode.style.display = "none";
            //alert(this.tagName);
        };
        deleteLink.innerText = "x";

        o.appendChild(deleteLink);
        //把当前画出的对象加入到画布中
        Rect.container.appendChild(o);
        //处理onmousemove事件
        Rect.container.onmousemove = Rect.move;
        //处理onmouseup事件
        Rect.container.onmouseup = Rect.end;
    },
    move: function (e) {
        let o = Rect.obj;
        //dx，dy是鼠标移动的距离
        let dx = Rect.getEvent(e).x - o.mouseBeginX;
        let dy = Rect.getEvent(e).y - o.mouseBeginY;
        //如果dx，dy <0,说明鼠标朝左上角移动，需要做特别的处理
        if (dx < 0) {
            o.style.left = Rect.getEvent(e).layerX + 'px';
        }
        if (dy < 0) {
            o.style.top = Rect.getEvent(e).clientY + 'px';
        }
        o.style.height = Math.abs(dy) + 'px';
        o.style.width = Math.abs(dx) + 'px';
        console.log(o.style.width + ',' + o.style.height);
    },
    end: function (e) {
        //onmouseup时释放onmousemove，onmouseup事件句柄
        Rect.container.onmousemove = null;
        Rect.container.onmouseup = null;
        Rect.obj = null;
        console.log(e);
    },
    //辅助方法，处理IE和FF不同的事件模型
    getEvent: function (e) {
        if (typeof e == 'undefined') {
            e = window.event;
        }
        //alert(e.x?e.x : e.layerX);
        if (typeof e.x == 'undefined') {
            e.x = e.layerX;
        }
        if (typeof e.y == 'undefined') {
            e.y = e.layerY;
        }
        return e;
    }
};

export default Rect;