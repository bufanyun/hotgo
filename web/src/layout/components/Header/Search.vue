<script setup lang="ts">
import { watchEffect, computed, nextTick, ref } from 'vue';
import { useRouter } from 'vue-router';
import useBoolean from "@/hooks/useBoolean";
import { useMagicKeys  } from "@vueuse/core";
import { useAsyncRouteStore } from '@/store/modules/asyncRoute';
import { adaModalWidth } from '@/utils/hotgo';
import { tree2FlatArray } from '@/utils/tree';
import { cloneDeep, debounce } from 'lodash-es';

const routeStore = useAsyncRouteStore();
const meun = routeStore.getMenus
const menusFlat = computed(()=>{
  const copyMeun = cloneDeep(meun)
  // 菜单打平
  return tree2FlatArray(copyMeun)
})
// 搜索值
const searchValue = ref('')

// 选中索引
const selectedIndex = ref<number>(0)

const { bool: showModal, setTrue: openModal, setFalse: closeModal, toggle: toggleModal } = useBoolean(false)

// 鼠标和键盘操作切换锁，防止鼠标和键盘操作冲突
const { bool: keyboardFlag, setTrue: setKeyboardTrue, setFalse: setKeyboardFalse } = useBoolean(false)

const { ctrl_k, arrowup, arrowdown, enter } = useMagicKeys({
  passive: false,
  onEventFired(e) {
    if (e.ctrlKey && e.key === 'k' && e.type === 'keydown')
      e.preventDefault()
  },
})

// 监听全局热键
watchEffect(() => {
  if (ctrl_k.value)
    toggleModal()
})


// 计算符合条件的菜单选项
const options = ref([])
const getFilterMenuOptions = debounce(()=> {
  selectedIndex.value = 0
  if (!searchValue.value){
    options.value = []
    return
  }
  
  const list = menusFlat.value.filter((item) => {
    const conditions = [
      item.name?.includes(searchValue.value),
      (item?.meta?.title || '')?.includes(searchValue.value),
      item.path?.includes(searchValue.value),
    ]
    return conditions.some(condition => condition)
  }).map((item) => {
    return {
      label: item?.meta?.title || item.name,
      icon: item.icon,
      value: item.path
    }
  })

  options.value = list
}, 200)

const router = useRouter()

const dialogWidth = computed(() => {
  return adaModalWidth(600);
});

// 关闭回调
function handleClose() {
  searchValue.value = ''
  selectedIndex.value = 0
  closeModal()
}

// 输入框改变，索引重置
function handleInputChange() {
  
  getFilterMenuOptions()
  
}

// 选择菜单选项
function handleSelect(value: string) {
  handleClose()
  router.push(value)
  nextTick(() => {
    searchValue.value = ''
  })
}

watchEffect(() => {
  // 没有打开弹窗或没有搜索结果时，不操作
  if (!showModal.value || !options.value.length)
    return

  // 设置键盘操作锁，设置后不会被动触发mouseover
  setKeyboardTrue()
  if (arrowup.value)
    handleArrowup()

  if (arrowdown.value)
    handleArrowdown()

  if (enter.value)
    handleEnter()
})

const scrollbarRef = ref()

// 上箭头操作
function handleArrowup() {
  if (selectedIndex.value === 0)
    selectedIndex.value = options.value.length - 1

  else
    selectedIndex.value--

  handleScroll(selectedIndex.value)
}

// 下箭头操作
function handleArrowdown() {
  if (selectedIndex.value === options.value.length - 1)
    selectedIndex.value = 0

  else
    selectedIndex.value++

  handleScroll(selectedIndex.value)
}

function handleScroll(currentIndex: number) {
  // 保持6个选项在可视区域内,6个后开始滚动
  const keepIndex = 5
  // 单个元素的高度，包括了元素的gap和容器的padding
  const elHeight = 70
  const distance = currentIndex * elHeight > keepIndex * elHeight ? currentIndex * elHeight - keepIndex * elHeight : 0
  scrollbarRef.value?.scrollTo({
    top: distance,
  })
}
// 回车键操作
function handleEnter() {
  const target = options.value[selectedIndex.value]
  if (target)
    handleSelect(target.value)
}

// 鼠标移入操作
function handleMouseEnter(index: number) {
  if (keyboardFlag.value)
    return

  selectedIndex.value = index
}
</script>

<template>
  <div class="flex items-center" @click="openModal" title="点击搜索菜单">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path fill="#000000" d="m19.485 20.154l-6.262-6.262q-.75.639-1.725.989t-1.96.35q-2.402 0-4.066-1.663T3.808 9.503T5.47 5.436t4.064-1.667t4.068 1.664T15.268 9.5q0 1.042-.369 2.017t-.97 1.668l6.262 6.261zM9.539 14.23q1.99 0 3.36-1.37t1.37-3.361t-1.37-3.36t-3.36-1.37t-3.361 1.37t-1.37 3.36t1.37 3.36t3.36 1.37"/></svg>
    <n-tag round size="small" class="font-mono cursor-pointer">
      Ctrl K
    </n-tag>
  </div>
  <n-modal
    v-model:show="showModal"
    class="fixed top-[60px] inset-x-0"
    size="small"
    preset="card"
    :segmented="{
      content: true,
      footer: true,
    }"
    :style="{
      width: dialogWidth,
    }"
    :closable="false"
    @after-leave="handleClose"
  >
    <template #header>
      <n-input v-model:value="searchValue" placeholder="搜索页面/路径" clearable size="large" @input="handleInputChange">
        <template #prefix>
          <n-icon>
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path fill="#6f6a6a" d="m19.485 20.154l-6.262-6.262q-.75.639-1.725.989t-1.96.35q-2.402 0-4.066-1.663T3.808 9.503T5.47 5.436t4.064-1.667t4.068 1.664T15.268 9.5q0 1.042-.369 2.017t-.97 1.668l6.262 6.261zM9.539 14.23q1.99 0 3.36-1.37t1.37-3.361t-1.37-3.36t-3.36-1.37t-3.361 1.37t-1.37 3.36t1.37 3.36t3.36 1.37"/></svg>
          </n-icon>
        </template>
      </n-input>
    </template>
    <n-scrollbar ref="scrollbarRef" class="h-[450px]">
      <ul
        v-if="options.length"
        class="flex flex-col gap-[8px] p-[8px] p-r-3"
      >
        <n-el
          v-for="(option, index) in options"
          :key="option.value" tag="li" role="option"
          class="cursor-pointer shadow h-[62px]"
          :class="{ 'text-[var(--base-color)] bg-[var(--primary-color-hover)]': index === selectedIndex }"
          @click="handleSelect(option.value)"
          @mouseenter="handleMouseEnter(index)"
          @mousemove="setKeyboardFalse"
        >
          <div class="grid grid-cols-[1fr_30px] h-full p-2">
            <span>{{ option.label }}</span>
            <svg class="row-span-2 place-self-center" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path fill="#6f6a6a" d="m13.292 12l-4.6-4.6l.708-.708L14.708 12L9.4 17.308l-.708-.708z"/></svg>
            <span class="op-70">{{ option.value }}</span>
          </div>
        </n-el>
      </ul>

      <n-empty v-else size="large" class="h-[450px] flex-center" />
    </n-scrollbar>

    <template #footer>
      <n-flex>
        <div class="flex items-center gap-1">
          <svg width="15" height="15" aria-label="Enter key" role="img"><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.2"><path d="M12 3.53088v3c0 1-1 2-2 2H4M7 11.53088l-3-3 3-3" /></g></svg>
          <span>选择</span>
        </div>
        <div class="flex items-center gap-1">
          <svg width="15" height="15" aria-label="Arrow down" role="img"><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.2"><path d="M7.5 3.5v8M10.5 8.5l-3 3-3-3" /></g></svg>
          <svg width="15" height="15" aria-label="Arrow up" role="img"><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.2"><path d="M7.5 11.5v-8M10.5 6.5l-3-3-3 3" /></g></svg>
          <span>切换</span>
        </div>
        <div class="flex items-center gap-1">
          <svg width="15" height="15" aria-label="Escape key" role="img"><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.2"><path d="M13.6167 8.936c-.1065.3583-.6883.962-1.4875.962-.7993 0-1.653-.9165-1.653-2.1258v-.5678c0-1.2548.7896-2.1016 1.653-2.1016.8634 0 1.3601.4778 1.4875 1.0724M9 6c-.1352-.4735-.7506-.9219-1.46-.8972-.7092.0246-1.344.57-1.344 1.2166s.4198.8812 1.3445.9805C8.465 7.3992 8.968 7.9337 9 8.5c.032.5663-.454 1.398-1.4595 1.398C6.6593 9.898 6 9 5.963 8.4851m-1.4748.5368c-.2635.5941-.8099.876-1.5443.876s-1.7073-.6248-1.7073-2.204v-.4603c0-1.0416.721-2.131 1.7073-2.131.9864 0 1.6425 1.031 1.5443 2.2492h-2.956" /></g></svg>
          <span>关闭</span>
        </div>
      </n-flex>
    </template>
  </n-modal>
</template>
