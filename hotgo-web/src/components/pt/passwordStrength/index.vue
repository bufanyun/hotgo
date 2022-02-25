<template>
  <div class="container">
    <a-row>
      <a-col :span="12">
        <a-form-model-item has-feedback :label="tip" :prop="prop">
          <a-input-password v-model="pwdee" id="inputValue" :placeholder="'请输入'+tip" />
        </a-form-model-item>
      </a-col>
      <a-col :span="10">
        <div class="input_span" style="width:240px;margin-top: 40px;margin-left: 15px;">
          <label>安全强度:</label>
          <span id="one">弱</span>
          <span id="two">中</span>
          <span id="three">强</span>
        </div>
      </a-col>
      <!-- <div id="font">
      <span>弱</span>
      <span>中</span>
      <span>强</span>
    </div> -->
    </a-row>
  </div>
</template>

<script>
export default {
  props: {
    tip: { type: String, default: '密码' },
    prop: String
  },
  data () {
    return { msgText: '', pwdee: '' }
  },
  methods: {
    checkStrong (sValue) {
      var modes = 0
      // 正则表达式验证符合要求的
      if (sValue.length < 1) return modes
      if (/\d/.test(sValue)) modes++ // 数字
      if (/[a-z]/.test(sValue)) modes++ // 小写
      if (/[A-Z]/.test(sValue)) modes++ // 大写
      if (/\W/.test(sValue)) modes++ // 特殊字符

      // 逻辑处理
      switch (modes) {
        case 1:
          return 1
        case 2:
          return 2
        case 3:
        case 4:
          return sValue.length < 4 ? 3 : 4
      }
      return modes
    }
  },
  components: {},
  watch: {
    pwdee (newname, oldname) {
      this.msgText = this.checkStrong(newname)
      if (this.msgText > 1 || this.msgText === 1) {
        document.getElementById('one').style.background = 'red'
      } else {
        document.getElementById('one').style.background = '#cccccc'
      }
      if (this.msgText > 2 || this.msgText === 2) {
        document.getElementById('two').style.background = 'orange'
      } else {
        document.getElementById('two').style.background = '#cccccc'
      }
      if (this.msgText === 4) {
        document.getElementById('three').style.background = '#00D1B2'
      } else {
        document.getElementById('three').style.background = '#cccccc'
      }
      this.$emit('input', newname)
    }
  }
}
</script>

<style scoped>
#inputValue {
  width: 240px;
  margin-left: 20px;
  padding-left: 10px;
  border-radius: 3px;
}
.input_span label{
  margin-right: 10px;
}
.input_span span {
  display: inline-block;
  width: 54px;
  height: 16px;
  background: #cccccc;
  line-height: 16px;
  margin-right: 2px;
  text-align: center;
  color: #ffffff;
}

#one {

}

#two {

}

#three {

}
#font span:nth-child(1) {
  color: red;
  margin-left: 80px;
}
#font span:nth-child(2) {
  color: orange;
  margin: 0 60px;
}
#font span:nth-child(3) {
  color: #00d1b2;
}
</style>
