<template>
  <a-modal
    :id="modalId"
    :ref="modalRefName"
    :centered="true"
    v-bind="$attrs"
    :body-style="ModalWidthAndHeight.bodyStyle"
    :style="ModalWidthAndHeight.modalStyle"
    :width="ModalWidthAndHeight.modalStyle.width"
    :closable="closeAble"
    v-on="$listeners"
    modalCutHeight:120
  >
    <template slot="title" v-if="isShowTitle">
      {{ modalTitle }}
      <div v-if="!sizeState" v-drag />
      <button
        v-if="adjustSize"
        :title="btnTitle"
        type="button"
        class="aidex-modal-size"
        :style="sizeBtnPosition"
        @click="fullScreen($event)"
      >
        <span class="aidex-modal-size-adjust">
          <a-icon :type="iconType()" />
        </span>
      </button>
    </template>
    <div class="modal-content">
      <slot name="content"></slot>
    </div>
    <template slot="footer">
      <slot name="footer" />
    </template>
  </a-modal>
</template>
<script>
import { Modal } from 'ant-design-vue'
import { appointModalWidthAndHeight } from '@/utils/pt/layout/baseMouldStyles'
let newSetModalStyle, newSetModalBodyStyle
export default {
  name: 'AntModal',
  components: {
    AModal: Modal
  },
  props: {
    // 弹窗标题
    modalTitle: {
      type: String,
      required: false,
      default: null
    },
    // 是否需要右上角全屏按钮
    adjustSize: {
      type: Boolean,
      required: false,
      default: true
    },
    isShowTitle: {
      type: Boolean,
      required: false,
      default: true
    },
    // 是否需要右上角关闭按钮
    closeAble: {
      type: Boolean,
      required: false,
      default: true
    },
    // 弹窗整体样式
    setModalStyle: {
      type: Object,
      required: false,
      default: null
    },
    // 弹窗中间body部分样式
    setModalBodyStyle: {
      type: Object,
      required: false,
      default: null
    },
    // 弹窗modal高度
    modalHeight: {
      type: String,
      required: false,
      default: '480'
    },
    // 弹窗modal宽度
    modalWidth: {
      type: String,
      required: false,
      default: '640'
    },
    // 弹窗modal的body的高度比整个弹窗modal少的高度
    modalCutHeight: {
      type: Number,
      default: 110
    }
  },
  data () {
    return {
      modalRefName: 'aModal',
      sizeState: 0, // 代表现在是正常状态，sizeState:1代表现在是放大状态
      moveEl: null,
      cutHeight: this.isShowTitle ? this.modalCutHeight : 0,
      ModalWidthAndHeight: appointModalWidthAndHeight(
        this.modalHeight,
        this.modalWidth,
        this.isShowTitle ? this.modalCutHeight : 0
      ),
      sizeBtnPosition: {
        right: '56px'
      },
      oldPosLeft: null,
      oldPosTop: null
    }
  },
  watch: {
    // sizeState: {
    //   immediate: true,
    //   handler(newV) {
    //     if (newV) {
    //       // let that = this
    //     }
    //   }
    // }
  },
  created () {
  },
  computed: {
    modalId () {
      return 'modalId' + this._uid
    },
    btnTitle () {
      if (!this.sizeState) {
        return '放大'
      } else {
        return '缩小'
      }
    }
  },
  mounted () {
    window.addEventListener('resize', () => {
      if (this.sizeState === 0) {
        this.ModalWidthAndHeight = appointModalWidthAndHeight(
          this.modalHeight,
          this.modalWidth,
          this.cutHeight
        )
      } else {
        this.ModalWidthAndHeight = appointModalWidthAndHeight('100%', '100%', this.cutHeight)
      }
    })
    if (this.setModalStyle != null && typeof (this.setModalStyle === Object)) {
      Object.assign(this.ModalWidthAndHeight.modalStyle, this.setModalStyle)
    }
    if (this.setModalBodyStyle != null && typeof (this.setModalBodyStyle === Object)) {
      Object.assign(this.ModalWidthAndHeight.bodyStyle, this.setModalBodyStyle)
    }
    if (!this.closeAble) {
      this.sizeBtnPosition.width = 0
    }
  },
  methods: {
    setMaxDiolog () {
      this.sizeState = 0
      this.fullScreen()
    },
    reduceScreen () {
      const moveEl = document.querySelector('#' + this.modalId + ' .ant-modal')
      this.ModalWidthAndHeight = appointModalWidthAndHeight(
        this.modalHeight,
        this.modalWidth,
        this.cutHeight
      )
      if (this.setModalStyle != null && typeof (this.setModalStyle === Object)) {
        Object.assign(this.ModalWidthAndHeight.modalStyle, this.setModalStyle)
      }
      if (this.setModalBodyStyle != null && typeof (this.setModalBodyStyle === Object)) {
        Object.assign(this.ModalWidthAndHeight.bodyStyle, this.setModalBodyStyle)
      }
      this.$nextTick(() => {
      moveEl.style.position = 'absolute'
      moveEl.style.left = this.oldPosLeft + 'px'
      moveEl.style.top = this.oldPosTop + 'px'

     /* moveEl.style.position = 'relative'
      moveEl.style.left = 'auto'
      moveEl.style.top = 'auto' */
      })
      this.sizeState = 0
    },
    fullScreen () {
      if (this.sizeState === 0) {
        const moveEl = document.querySelector('#' + this.modalId + ' .ant-modal')
        const modalPosition = this.getModalPosition(moveEl)
        this.oldPosLeft = modalPosition.left
        this.oldPosTop = modalPosition.top
        moveEl.style.left = '0px'
        moveEl.style.top = '0px'
        console.log(this.oldPosLeft)
        console.log(this.oldPosTop)
        this.ModalWidthAndHeight = appointModalWidthAndHeight('100%', '100%', this.cutHeight)

        if (this.setModalStyle != null && typeof (this.setModalStyle === Object)) {
          newSetModalStyle = JSON.parse(JSON.stringify(this.setModalStyle))
          if (newSetModalStyle.height) {
            newSetModalStyle.height = this.ModalWidthAndHeight.modalStyle.height
          }
          if (newSetModalStyle.width) {
            newSetModalStyle.width = this.ModalWidthAndHeight.modalStyle.width
          }
          Object.assign(this.ModalWidthAndHeight.modalStyle, newSetModalStyle)
        }
        if (this.setModalBodyStyle != null && typeof (this.setModalBodyStyle === Object)) {
          newSetModalBodyStyle = JSON.parse(JSON.stringify(this.setModalBodyStyle))
          if (newSetModalBodyStyle.height) {
            newSetModalBodyStyle.height = this.ModalWidthAndHeight.bodyStyle.height
          }
          if (newSetModalBodyStyle.width) {
            newSetModalBodyStyle.width = this.ModalWidthAndHeight.bodyStyle.width
          }
          Object.assign(this.ModalWidthAndHeight.bodyStyle, newSetModalBodyStyle)
        }
        // this.$nextTick(() => {

        //   moveEl.style.left = '7.5px'
        //   moveEl.style.top = '0px'
        // })
        this.sizeState = 1
        document.querySelector('#' + this.modalId + ' .ant-modal-wrap').style.overflow = 'hidden'
      } else {
        this.reduceScreen()
        document.querySelector('#' + this.modalId + ' .ant-modal-wrap').style.overflow = 'auto'
      }
    },
    iconType () {
      if (this.sizeState === 0) {
        return 'fullscreen'
      } else {
        return `fullscreen-exit`
      }
    },
    getModalPosition (el) {
      return el.getBoundingClientRect()
    }
  }
}
</script>

<style lang="less">
.aidex-modal-size {
  position: absolute;
  top: 0;
  right: 56px;
  z-index: 10;
  padding: 0;
  color: rgba(0, 0, 0, 0.45);
  font-weight: 700;
  line-height: 1;
  text-decoration: none;
  background: transparent;
  border: 0;
  outline: 0;
  cursor: pointer;
  transition: color 0.3s;
}
.aidex-modal-size-adjust {
  display: block;
  width: 56px;
  height: 56px;
  font-size: 16px;
  font-style: normal;
  line-height: 56px;
  text-align: center;
  text-transform: none;
  text-rendering: auto;
}
.aidex-modal-size:hover {
  color: rgba(0, 0, 0, 0.75);
  text-decoration: none;
}
.ant-modal-body{
   .modal-content{
     .advanced-table{
       min-height: auto;
       .ant-table-body{
          min-height: auto;
       }
     }
   }
   }
</style>
