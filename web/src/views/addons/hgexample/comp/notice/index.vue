<template>
  <div>
    <n-grid :cols="2" :x-gap="10" :y-gap="10">
      <n-gi>
        <n-card
          title="警告信息alert"
          :content-style="{ padding: '10px' }"
          :header-style="{ padding: '10px' }"
        >
          <n-space vertical :size="12">
            <n-alert title="Info 类型" type="info"> 万事开头难，只能一直努力！！！ </n-alert>
            <n-alert title="Success 类型" type="success">
              成功就是眼前，过了黑夜就是白天，加油~~
            </n-alert>
            <n-alert title="Warning 类型" type="warning">
              想想昨天，看看今天，望望明天；一切都会好起来
            </n-alert>
            <n-alert title="Error 类型" type="error"> 有人欺负人，请不犹豫，请还回去 </n-alert>
            <n-alert title="可以关掉" type="info" closable> Gee it's good to be back home </n-alert>
            <n-alert title="没有边框" type="info" :bordered="false">
              Gee it's good to be back home
            </n-alert>
            <n-alert title="自定义图标">
              <template #icon>
                <n-icon>
                  <BalloonOutline />
                </n-icon>
              </template>
              Well the Ukraine girls really knock me out<br />
              They leave the West behind<br />
              And Moscow girls make me sing and shout<br />
              That Georgia's always on my mind<br />
              Aw come on!
            </n-alert>

            <n-alert :show-icon="false" title="没有图标">
              Yeah I'm back in the U.S.S.R.<br />
              You don't know how lucky you are boys
            </n-alert>
          </n-space>
        </n-card>
      </n-gi>
      <n-gi>
        <n-card
          title="对话框Dialog"
          :content-style="{ padding: '10px' }"
          :header-style="{ padding: '10px' }"
        >
          <div class="padding">
            <n-space>
              <n-button size="small" type="primary" @click="openConfirm('warning')">警告</n-button>
              <n-button size="small" type="warning" @click="openConfirm('success')">成功</n-button>
              <n-button size="small" type="info" @click="openConfirm('error')">错误</n-button>
            </n-space>
          </div>
        </n-card>
        <n-card
          class="mt-2"
          title="信息Message"
          :content-style="{ padding: '10px' }"
          :header-style="{ padding: '10px' }"
        >
          <n-space>
            <n-button type="info" @click="openMessage('info')"> 信息 </n-button>
            <n-button type="error" @click="openMessage('error')"> 错误 </n-button>
            <n-button type="warning" @click="openMessage('warning')"> 警告 </n-button>
            <n-button type="success" @click="openMessage('success')"> 成功 </n-button>
            <n-button type="primary" @click="openMessage('loading')"> 加载中 </n-button>
          </n-space>
        </n-card>
        <n-card
          class="mt-2"
          title="通知Notification"
          :content-style="{ padding: '10px' }"
          :header-style="{ padding: '10px' }"
        >
          <n-space>
            <n-button type="info" @click="openNotification('info')"> 信息 </n-button>
            <n-button type="error" @click="openNotification('error')"> 错误 </n-button>
            <n-button type="warning" @click="openNotification('warning')"> 警告 </n-button>
            <n-button type="success" @click="openNotification('success')"> 成功 </n-button>
          </n-space>
        </n-card>

        <n-card
          class="mt-2"
          title="弹出确认Popconfirm"
          :content-style="{ padding: '10px' }"
          :header-style="{ padding: '10px' }"
        >
          <n-space>
            <n-popconfirm
              @positive-click="handlePositiveClick"
              @negative-click="handleNegativeClick"
            >
              <template #trigger>
                <n-button>引用</n-button>
              </template>
              一切都将一去杳然，任何人都无法将其捕获。
            </n-popconfirm>
            <n-popconfirm :negative-text="null" @positive-click="handlePositiveClick">
              <template #trigger>
                <n-button>只有确定</n-button>
              </template>
              一切都将一去杳然，任何人都无法将其捕获。
            </n-popconfirm>
            <n-popconfirm :positive-text="null" @negative-click="handleNegativeClick">
              <template #trigger>
                <n-button>只有取消</n-button>
              </template>
              一切都将一去杳然，任何人都无法将其捕获。
            </n-popconfirm>
            <n-popconfirm :positive-text="null" :negative-text="null">
              <template #trigger>
                <n-button>什么也没有</n-button>
              </template>
              一切都将一去杳然，任何人都无法将其捕获。
            </n-popconfirm>
            <n-popconfirm>
              <template #trigger>
                <n-button>自定义 action</n-button>
              </template>
              <template #action> 自定义 action </template>
              一切都将一去杳然，任何人都无法将其捕获。
            </n-popconfirm>
          </n-space>
        </n-card>

        <n-card
          class="mt-2"
          title="弹出信息Popover"
          :content-style="{ padding: '10px' }"
          :header-style="{ padding: '10px' }"
        >
          <n-space>
            <n-popover trigger="hover">
              <template #trigger>
                <n-button>悬浮</n-button>
              </template>
              <span>I wish they all could be California girls</span>
            </n-popover>
            <n-popover trigger="hover" :keep-alive-on-hover="false">
              <template #trigger>
                <n-button>悬浮（忽略主体）</n-button>
              </template>
              <span>I wish they all could be California girls</span>
            </n-popover>
            <n-popover trigger="click">
              <template #trigger>
                <n-button>点击</n-button>
              </template>
              <span>I wish they all could be California girls</span>
            </n-popover>
            <n-popover trigger="focus">
              <template #trigger>
                <n-button>聚焦</n-button>
              </template>
              <span>I wish they all could be California girls</span>
            </n-popover>
            <n-popover trigger="manual" :show="showPopover">
              <template #trigger>
                <n-button @click="showPopover = !showPopover"> 手动 </n-button>
              </template>
              <span>I wish they all could be California girls</span>
            </n-popover>
            <n-popover trigger="hover" :delay="500" :duration="500">
              <template #trigger>
                <n-button>延迟 500ms, 持续 500ms</n-button>
              </template>
              <span>
                Lately did you ever feel the pain In the morning rain as it soaks it to the bone
              </span>
            </n-popover>
          </n-space>
        </n-card>
        <n-card
          class="mt-2"
          title="弹出提示Tooltip"
          :content-style="{ padding: '10px' }"
          :header-style="{ padding: '10px' }"
        >
          <n-space>
            <n-tooltip trigger="hover">
              <template #trigger>
                <n-button> 鸭子 </n-button>
              </template>
              如果它长得像鸭子，走起来像鸭子，叫起来也像鸭子，那它一定是个鸭子。
            </n-tooltip>
            <n-tooltip placement="bottom" trigger="hover">
              <template #trigger>
                <n-button> 悬浮 </n-button>
              </template>
              <span> I wish they all could be California girls </span>
            </n-tooltip>
            <n-tooltip placement="bottom" trigger="click">
              <template #trigger>
                <n-button> 点击 </n-button>
              </template>
              <span> I wish they all could be California girls </span>
            </n-tooltip>
            <n-tooltip :show="showPopover" placement="bottom">
              <template #trigger>
                <n-button @click="showPopover = !showPopover"> 手动 </n-button>
              </template>
              <span> I wish they all could be California girls </span>
            </n-tooltip>
            <n-tooltip :show-arrow="false" trigger="hover">
              <template #trigger>
                <n-button>默认有箭头</n-button>
              </template>
              和 Popover 一样
            </n-tooltip>
          </n-space>
        </n-card>
      </n-gi>
    </n-grid>
  </div>
</template>

<script lang="ts" setup>
  import { ref } from 'vue';
  import { useDialog, useMessage, useNotification } from 'naive-ui';
  import { BalloonOutline } from '@vicons/ionicons5';

  const dialog = useDialog();
  const message = useMessage();
  const notification = useNotification();
  const showPopover = ref(false);

  function openConfirm(type: string) {
    switch (type) {
      case 'warning':
        dialog.warning({
          title: '警告',
          content: '你确定要离开我？',
          positiveText: '确定',
          negativeText: '不确定',
          onPositiveClick: () => {
            message.success('确定');
          },
          onNegativeClick: () => {
            message.error('不确定');
          },
        });
        break;
      case 'success':
        dialog.success({
          title: '成功',
          content: '小哥哥你太棒了~~',
          positiveText: '哇哦~',
        });
        break;
      case 'error':
        dialog.error({
          title: '错误',
          content: '知道错了吗，下次改了',
          positiveText: '快哭了',
          onPositiveClick: () => {
            message.success('我知道了');
          },
        });
        break;
    }
  }

  function openMessage(type: string) {
    switch (type) {
      case 'info':
        message.info('在心碎中认清遗憾 生命漫长也短暂跳动心脏长出藤蔓');
        break;
      case 'error':
        message.error('愿为险而战 跌入灰暗坠入深渊');
        break;
      case 'warning':
        message.warning('沾满泥土的脸 没有神的光环');
        break;
      case 'success':
        message.success('握紧手中的平凡 此心自称无憾生命的火已点燃');
        break;
      case 'loading':
        message.loading('有一天也许会走远 也许还能再相见');
        break;
    }
  }

  function openNotification(type: string) {
    switch (type) {
      case 'info':
        notification.info({
          content: '在心碎中认清遗憾 生命漫长也短暂跳动心脏长出藤蔓',
          meta: '只要平凡',
        });
        break;
      case 'error':
        notification.error({
          content: '愿为险而战 跌入灰暗坠入深渊',
          meta: '只要平凡',
        });
        break;
      case 'warning':
        notification.warning({
          content: '沾满泥土的脸 没有神的光环',
          meta: '只要平凡',
        });
        break;
      case 'success':
        notification.success({
          content: '握紧手中的平凡 此心自称无憾生命的火已点燃',
          meta: '只要平凡',
        });
        break;
    }
  }

  function handlePositiveClick() {
    message.info('是的');
  }

  function handleNegativeClick() {
    message.info('并不');
  }
</script>
