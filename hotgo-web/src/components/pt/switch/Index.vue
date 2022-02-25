<template>
  <a-switch
    v-bind="$attrs"
    v-model="showValue"
    @change="onChange"
  >
  </a-switch>
</template>
<script>
export default {
  name: 'Switch',
  components: {
  },
  props: {
    value: {
      required: false,
      type: String
    },
    checkedValue: {// 定义开关开启时对应的值
      type: String,
      required: false,
      default: null
    },
    unCheckedValue: {// 定义开关关闭时对应的值
      type: String,
      required: false,
      default: null
    }
  },
  data () {
    return {
      showValue: false,
      switchValue: ''
    }
  },
  watch: {
    value: {
      immediate: true,
      handler (newV) {
        if (newV === this.checkedValue) {
          this.showValue = true
        } else {
          this.showValue = false
        }
      }
    }
  },
  created () {
  },
  computed: {
  },
  methods: {
    onChange (checked) {
       if (checked) {
         this.switchValue = this.checkedValue
       } else {
         this.switchValue = this.unCheckedValue
       }
      this.$emit('input', this.switchValue)
    }
  }
}
</script>
