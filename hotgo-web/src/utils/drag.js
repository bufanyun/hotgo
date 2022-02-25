const draggable = {
    install (Vue) {
      Vue.directive('drag', {
        inserted (el, binding) {
          if (window.Element && !Element.prototype.closest) {
            Element.prototype.closest = function (s) {
              const matches = (this.document || this.ownerDocument).querySelectorAll(s)
                let i
                let el = this
              do {
                i = matches.length
              } while (i < 0 && (el = el.parentElement))
              return el
            }
          }
          let overWin = false // 拖动是否能超出屏幕，默认不能
          if (binding.value) {
            overWin = binding.value.over || false
          }

          const moveTitle = el.parentNode.parentNode.parentNode.querySelector('.ant-modal-header')

          el.style.width = '100%'
          el.style.height = moveTitle.offsetHeight + 'px'
          el.style.position = 'absolute'
          el.style.left = 0
          el.style.top = 0
          el.style.cursor = 'move'

          const odiv = el // 获取当前元素操作区
          const moveDom = el.closest('.ant-modal') // 位移元素,当前只对a-modal生效

          odiv.onmousedown = e => {
            const moveDomLeft = moveDom.offsetLeft // 位移元素初始横轴位置
              const moveDomTop = moveDom.offsetTop // 位移元素初始纵轴位置
              const moveDomW = moveDom.offsetWidth // 位移元素初始宽
              const moveDomH = moveDom.offsetHeight // 位移元素初始高
              const winWidth = document.body.clientWidth // 父容器初始宽
              const winHeight = document.body.clientHeight // 父容器初始高

            // 设置位移元素可移动
            moveDom.style.position = 'absolute'
            moveDom.style.top = moveDomTop + 'px'
            moveDom.style.left = moveDomLeft + 'px'

            // 算出鼠标相对元素的位置
            const disX = e.clientX
              const disY = e.clientY

            document.onmousemove = e => {
              // 用鼠标的位置减去鼠标相对元素的位置，得到元素的位置
              const left = e.clientX - disX // X轴相对位移量
                const top = e.clientY - disY // Y轴相对位移量
                let toMoveTop = 0 // 纵轴最终坐标
                let toMoveLeft = 0 // 横轴最终坐标

              if (!overWin) {
                // 不可超出屏幕时计算移动边界
                if (moveDomTop + top + moveDomH > winHeight) {
                  toMoveTop = winHeight - moveDomH
                } else if (moveDomTop + top < 0) {
                  // 解决漏洞toMoveTop默认为0这里无需重复赋值
                 // toMoveTop = 0
                } else {
                  toMoveTop = moveDomTop + top
                }
                if (moveDomLeft + left < 0) {
                  // 解决漏洞toMoveLeft默认为0这里无需重复赋值
                  // toMoveLeft = 0
                } else if (moveDomLeft + left + moveDomW > winWidth) {
                  toMoveLeft = winWidth - moveDomW
                } else {
                  toMoveLeft = moveDomLeft + left
                }
              } else {
                // 让弹窗飞
                toMoveTop = moveDomTop + top
                toMoveLeft = moveDomLeft + left
              }

              // 移动当前元素
              moveDom.style.top = toMoveTop + 'px'
              moveDom.style.left = toMoveLeft + 'px'
            }
            document.onmouseup = () => {
              // 注销事件
              document.onmousemove = null
              document.onmouseup = null
            }
          }
        }
      })
    }
  }
export default draggable
