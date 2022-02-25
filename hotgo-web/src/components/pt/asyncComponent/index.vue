<template>
  <transition
    v-if="transition && keepAlive"
    mode="out-in"
    :enter-class="transitionClass.enter"
    :enter-active-class="transitionClass.enterActive"
    :enter-to-class="transitionClass.enterTo"
    :leave-class="transitionClass.leave"
    :leave-active-class="transitionClass.leaveActive"
    :leave-to-class="transitionClass.leaveTo"
  >
    <keep-alive>
      <component :is="componentName" v-bind="$attrs" v-on="$listeners" />
    </keep-alive>
  </transition>
  <transition
    v-else-if="transition && !keepAlive"
    mode="out-in"
    :enter-class="transitionClass.enter"
    :enter-active-class="transitionClass.enterActive"
    :enter-to-class="transitionClass.enterTo"
    :leave-class="transitionClass.leave"
    :leave-active-class="transitionClass.leaveActive"
    :leave-to-class="transitionClass.leaveTo"
  >
    <component :is="componentName" v-bind="$attrs" v-on="$listeners" />
  </transition>
  <keep-alive v-else-if="!transition && keepAlive">
    <component :is="componentName" v-bind="$attrs" v-on="$listeners" />
  </keep-alive>
  <component v-else :is="componentName" v-bind="$attrs" v-on="$listeners" />
</template>
<script>
import asyncLoading from './Loading'
import asyncError from './Error'
export default {
  name: 'AsyncComponent',
  inheritAttrs: false,
  props: {
    // 需要异步加载的组件路径
    path: {
      type: String,
      required: true
    },
    // 是否需要保持之前加载的状态
    keepAlive: {
      type: Boolean,
      default: false
    },
    // 是否添加组件进入/离开过渡
    transition: {
      type: Boolean,
      default: false
    },
    // 组件显示延时时间
    delay: {
      type: Number,
      default: 200
    },
    // 组件加载超时时间
    timeout: {
      type: Number,
      default: 3000
    },
    // 过渡样式
    transitionClass: {
      type: Object,
      default () {
        return {
          enter: 'enter',
          enterTo: 'enter-to',
          enterActive: 'enter-active',
          leave: 'leave',
          leaveTo: 'leave-to',
          leaveActive: 'leave-active'
        }
      }
    }
  },
  data () {
    return {
      componentName: () => ({
        component: import(`@/${this.path}.vue`),
        loading: asyncLoading,
        error: asyncError,
        delay: this.delay,
        timeout: this.timeout
      })
    }
  },
  watch: {
    path () {
      this.componentName = () => ({
        component: import(`@/${this.path}.vue`),
        loading: asyncLoading,
        error: asyncError,
        delay: this.delay,
        timeout: this.timeout
      })
    }
  }
}
</script>
<style scoped>
.enter,
.leave-to {
  opacity: 0;
}
.enter-to,
.leave {
  opacity: 1;
}
.enter-active,
.leave-active {
  transition: opacity 0.5s;
}
</style>
