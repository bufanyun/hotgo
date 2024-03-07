## 表单组件

目录

- 文本输入 Input
- 数字输入 Input Number
- 文本域 InputTextarea
- 富文本 InputEditor
- 动态键值对 InputDynamic
- 日期选择 Date(Y-M-D)
- 日期范围选择 DateRange
- 时间选择 Time(Y-M-D H:i:s)
- 时间范围选择 TimeRange 
- 单选按钮 Radio
- 复选框 Checkbox
- 单选下拉框 Select
- 多选下拉框 SelectMultiple
- 树型选择 Tree Select
- 单图上传 UploadImage
- 多图上传 UploadImage
- 单文件上传 UploadFile
- 多文件上传 UploadFile
- 文件选择器 FileChooser
- 大文件上传 MultipartUpload
- 开关 Switch
- 评分 Rate
- 省市区选择器 CitySelector
- 图标选择器 IconSelector

### 文本输入 Input

```vue
<template>
<n-input v-model:value="value" type="text" placeholder="基本的 Input" />
</template>

<script lang="ts" setup>
import { ref } from 'vue'
const value = ref(null);
</script>
```

### 数字输入 Input Number

```vue
<template>
  <n-input-number v-model:value="value" clearable />
</template>

<script lang="ts" setup>
import { ref } from 'vue'
const value = ref(null);
</script>
```

### 文本域 InputTextarea

```vue
<template>
    <n-input
      v-model:value="value"
      type="textarea"
      placeholder="基本的 Textarea"
    />
</template>

<script lang="ts" setup>
import { ref } from 'vue'
const value = ref(null);
</script>
```

### 富文本 InputEditor

```vue
<template>
  <Editor style="height: 450px" v-model:value="value" />
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import Editor from '@/components/Editor/editor.vue';
const value = ref(null);
</script>
```

### 动态键值对 InputDynamic

```vue
<template>
  <n-dynamic-input
      v-model:value="value"
      preset="pair"
      key-placeholder="键名"
      value-placeholder="键值"
  />
</template>

<script lang="ts" setup>
import { ref } from 'vue'
const value = ref(null);
</script>
```

### 日期选择 Date(Y-M-D)

```vue
<template>
  <DatePicker v-model:formValue="value" type="date" />
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import DatePicker from '@/components/DatePicker/datePicker.vue';
const value = ref(null);
</script>
```

### 日期范围选择 DateRange

```vue
<template>
  <DatePicker
      v-model:startValue="startValue"
      v-model:endValue="endValue"
      type="datetimerange"
  />
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import DatePicker from '@/components/DatePicker/datePicker.vue';
const startValue = ref(null);
const endValue = ref(null);
</script>
```

### 时间选择 Time(Y-M-D H:i:s)

```vue
<template>
  <n-time-picker :default-formatted-value="value" />
</template>

<script lang="ts" setup>
import { ref } from 'vue'
const value = ref(null);
</script>
```

### 时间范围选择 TimeRange

```vue
<template>
  <template>
    <n-space>
      <n-time-picker :default-value="startValue" />
      <n-time-picker :default-value="endValue" />
    </n-space>
  </template>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
const startValue = ref(null);
const endValue = ref(null);
</script>
```

### 单选按钮 Radio

```vue
<template>
  <n-space vertical>
    <n-radio-group v-model:value="value" name="radiobuttongroup1">
      <n-radio-button
          v-for="song in songs"
          :key="song.value"
          :value="song.value"
          :disabled="
          (song.label === 'Live Forever' && disabled1) ||
            (song.label === 'Shakermaker' && disabled2)
        "
          :label="song.label"
      />
    </n-radio-group>
    <n-space>
      <n-checkbox v-model:checked="disabled2" style="margin-right: 12px">
        禁用 Shakemaker
      </n-checkbox>
      <n-checkbox v-model:checked="disabled1">
        禁用 Live Forever
      </n-checkbox>
    </n-space>
  </n-space>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'

export default defineComponent({
  setup () {
    return {
      value: ref(null),
      disabled2: ref(false),
      disabled1: ref(false),
      songs: [
        {
          value: "Rock'n'Roll Star",
          label: "Rock'n'Roll Star"
        },
        {
          value: 'Shakermaker',
          label: 'Shakermaker'
        },
        {
          value: 'Live Forever',
          label: 'Live Forever'
        },
        {
          value: 'Up in the Sky',
          label: 'Up in the Sky'
        },
        {
          value: '...',
          label: '...'
        }
      ].map((s) => {
        s.value = s.value.toLowerCase()
        return s
      })
    }
  }
})
</script>
```

### 复选框 Checkbox
```vue
<template>
  <n-space item-style="display: flex;" align="center">
    <n-checkbox v-model:checked="value">
      复选框
    </n-checkbox>
    <n-checkbox v-model:checked="value" />
    <n-checkbox v-model:checked="value" :disabled="disabled">
      复选框
    </n-checkbox>
    <n-button size="small" @click="disabled = !disabled">
      禁用
    </n-button>
  </n-space>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'

export default defineComponent({
  setup () {
    return {
      value: ref(false),
      disabled: ref(true)
    }
  }
})
</script>
```


### 单选下拉框 Select
```vue
<template>
  <n-space vertical>
    <n-select v-model:value="value" :options="options" />
    <n-select v-model:value="value" disabled :options="options" />
  </n-space>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'

export default defineComponent({
  setup () {
    return {
      value: ref(null),
      options: [
        {
          label: "Everybody's Got Something to Hide Except Me and My Monkey",
          value: 'song0',
          disabled: true
        },
        {
          label: 'Drive My Car',
          value: 'song1'
        },
        {
          label: 'Norwegian Wood',
          value: 'song2'
        },
        {
          label: "You Won't See",
          value: 'song3',
          disabled: true
        },
        {
          label: 'Nowhere Man',
          value: 'song4'
        },
        {
          label: 'Think For Yourself',
          value: 'song5'
        },
        {
          label: 'The Word',
          value: 'song6'
        },
        {
          label: 'Michelle',
          value: 'song7',
          disabled: true
        },
        {
          label: 'What goes on',
          value: 'song8'
        },
        {
          label: 'Girl',
          value: 'song9'
        },
        {
          label: "I'm looking through you",
          value: 'song10'
        },
        {
          label: 'In My Life',
          value: 'song11'
        },
        {
          label: 'Wait',
          value: 'song12'
        }
      ]
    }
  }
})
</script>
```

### 多选下拉框 SelectMultiple
```vue
<template>
  <n-space vertical>
    <n-select v-model:value="value" multiple :options="options" />
    <n-select v-model:value="value" multiple disabled :options="options" />
  </n-space>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'

export default defineComponent({
  setup () {
    return {
      value: ref(['song3']),
      options: [
        {
          label: "Everybody's Got Something to Hide Except Me and My Monkey",
          value: 'song0',
          disabled: true
        },
        {
          label: 'Drive My Car',
          value: 'song1'
        },
        {
          label: 'Norwegian Wood',
          value: 'song2'
        },
        {
          label: "You Won't See",
          value: 'song3',
          disabled: true
        },
        {
          label: 'Nowhere Man',
          value: 'song4'
        },
        {
          label: 'Think For Yourself',
          value: 'song5'
        },
        {
          label: 'The Word',
          value: 'song6'
        },
        {
          label: 'Michelle',
          value: 'song7',
          disabled: true
        },
        {
          label: 'What goes on',
          value: 'song8'
        },
        {
          label: 'Girl',
          value: 'song9'
        },
        {
          label: "I'm looking through you",
          value: 'song10'
        },
        {
          label: 'In My Life',
          value: 'song11'
        },
        {
          label: 'Wait',
          value: 'song12'
        }
      ]
    }
  }
})
</script>
```

### 树型选择 Tree Select
```vue
<template>
  <n-tree-select
    :options="options"
    default-value="Drive My Car"
    @update:value="handleUpdateValue"
  />
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { TreeSelectOption } from 'naive-ui'

export default defineComponent({
  setup () {
    return {
      handleUpdateValue (
        value: string | number | Array<string | number> | null,
        option: TreeSelectOption | null | Array<TreeSelectOption | null>
      ) {
        console.log(value, option)
      },
      options: [
        {
          label: 'Rubber Soul',
          key: 'Rubber Soul',
          children: [
            {
              label:
                "Everybody's Got Something to Hide Except Me and My Monkey",
              key: "Everybody's Got Something to Hide Except Me and My Monkey"
            },
            {
              label: 'Drive My Car',
              key: 'Drive My Car',
              disabled: true
            },
            {
              label: 'Norwegian Wood',
              key: 'Norwegian Wood'
            },
            {
              label: "You Won't See",
              key: "You Won't See",
              disabled: true
            },
            {
              label: 'Nowhere Man',
              key: 'Nowhere Man'
            },
            {
              label: 'Think For Yourself',
              key: 'Think For Yourself'
            },
            {
              label: 'The Word',
              key: 'The Word'
            },
            {
              label: 'Michelle',
              key: 'Michelle',
              disabled: true
            },
            {
              label: 'What goes on',
              key: 'What goes on'
            },
            {
              label: 'Girl',
              key: 'Girl'
            },
            {
              label: "I'm looking through you",
              key: "I'm looking through you"
            },
            {
              label: 'In My Life',
              key: 'In My Life'
            },
            {
              label: 'Wait',
              key: 'Wait'
            }
          ]
        },
        {
          label: 'Let It Be',
          key: 'Let It Be Album',
          children: [
            {
              label: 'Two Of Us',
              key: 'Two Of Us'
            },
            {
              label: 'Dig A Pony',
              key: 'Dig A Pony'
            },
            {
              label: 'Across The Universe',
              key: 'Across The Universe'
            },
            {
              label: 'I Me Mine',
              key: 'I Me Mine'
            },
            {
              label: 'Dig It',
              key: 'Dig It'
            },
            {
              label: 'Let It Be',
              key: 'Let It Be'
            },
            {
              label: 'Maggie Mae',
              key: 'Maggie Mae'
            },
            {
              label: "I've Got A Feeling",
              key: "I've Got A Feeling"
            },
            {
              label: 'One After 909',
              key: 'One After 909'
            },
            {
              label: 'The Long And Winding Road',
              key: 'The Long And Winding Road'
            },
            {
              label: 'For You Blue',
              key: 'For You Blue'
            },
            {
              label: 'Get Back',
              key: 'Get Back'
            }
          ]
        }
      ]
    }
  }
})
</script><template>
  <n-tree-select
    :options="options"
    default-value="Drive My Car"
    @update:value="handleUpdateValue"
  />
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { TreeSelectOption } from 'naive-ui'

export default defineComponent({
  setup () {
    return {
      handleUpdateValue (
        value: string | number | Array<string | number> | null,
        option: TreeSelectOption | null | Array<TreeSelectOption | null>
      ) {
        console.log(value, option)
      },
      options: [
        {
          label: 'Rubber Soul',
          key: 'Rubber Soul',
          children: [
            {
              label:
                "Everybody's Got Something to Hide Except Me and My Monkey",
              key: "Everybody's Got Something to Hide Except Me and My Monkey"
            },
            {
              label: 'Drive My Car',
              key: 'Drive My Car',
              disabled: true
            },
            {
              label: 'Norwegian Wood',
              key: 'Norwegian Wood'
            },
            {
              label: "You Won't See",
              key: "You Won't See",
              disabled: true
            },
            {
              label: 'Nowhere Man',
              key: 'Nowhere Man'
            },
            {
              label: 'Think For Yourself',
              key: 'Think For Yourself'
            },
            {
              label: 'The Word',
              key: 'The Word'
            },
            {
              label: 'Michelle',
              key: 'Michelle',
              disabled: true
            },
            {
              label: 'What goes on',
              key: 'What goes on'
            },
            {
              label: 'Girl',
              key: 'Girl'
            },
            {
              label: "I'm looking through you",
              key: "I'm looking through you"
            },
            {
              label: 'In My Life',
              key: 'In My Life'
            },
            {
              label: 'Wait',
              key: 'Wait'
            }
          ]
        },
        {
          label: 'Let It Be',
          key: 'Let It Be Album',
          children: [
            {
              label: 'Two Of Us',
              key: 'Two Of Us'
            },
            {
              label: 'Dig A Pony',
              key: 'Dig A Pony'
            },
            {
              label: 'Across The Universe',
              key: 'Across The Universe'
            },
            {
              label: 'I Me Mine',
              key: 'I Me Mine'
            },
            {
              label: 'Dig It',
              key: 'Dig It'
            },
            {
              label: 'Let It Be',
              key: 'Let It Be'
            },
            {
              label: 'Maggie Mae',
              key: 'Maggie Mae'
            },
            {
              label: "I've Got A Feeling",
              key: "I've Got A Feeling"
            },
            {
              label: 'One After 909',
              key: 'One After 909'
            },
            {
              label: 'The Long And Winding Road',
              key: 'The Long And Winding Road'
            },
            {
              label: 'For You Blue',
              key: 'For You Blue'
            },
            {
              label: 'Get Back',
              key: 'Get Back'
            }
          ]
        }
      ]
    }
  }
})
</script>
```


### 单图上传 UploadImage
```vue
<template>
  <UploadImage :maxNumber="1" v-model:value="value" />
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import UploadImage from '@/components/Upload/uploadImage.vue';
const value = ref(null);
</script>
```

### 多图上传 UploadImage
```vue
<template>
  <UploadImage :maxNumber="10" v-model:value="value" />
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import UploadImage from '@/components/Upload/uploadImage.vue';
const value = ref(null);
</script>
```

### 单文件上传 UploadFile
```vue
<template>
  <UploadFile :maxNumber="1" v-model:value="value" />
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import UploadFile from '@/components/Upload/uploadFile.vue';
const value = ref(null);
</script>
```

### 多文件上传 UploadFile
```vue
<template>
  <UploadFile :maxNumber="10" v-model:value="value" />
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import UploadFile from '@/components/Upload/uploadFile.vue';
const value = ref(null);
</script>
```

### 文件选择器 FileChooser
- 基础用法
```vue
<template>
  <FileChooser v-model:value="value" />
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import FileChooser from '@/components/FileChooser/index.vue';
const value = ref(null);
</script>
```

- 指定fileType，支持多种选择器类型，默认情况是全部都可以选择
```ts
type FileType = 'image' | 'doc' | 'audio' | 'video' | 'zip' | 'other' | 'default';
```

- 图片选择器
```vue
<FileChooser v-model:value="value" fileType="image" />
```

- 多选支持，指定`maxNumber`多选数量
```vue
<FileChooser v-model:value="value" :maxNumber="10" fileType="image" />
```

### 大文件上传 MultipartUpload
- 基础用法
```vue
<template>
  <MultipartUpload ref="multipartUploadRef" @onFinish="handleFinishCall" />
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import MultipartUpload from '@/components/Upload/multipartUpload.vue';
import { Attachment } from '@/components/FileChooser/src/model';
const multipartUploadRef = ref();

// 打开上传Modal
function handleMultipartUpload() {
  multipartUploadRef.value.openModal();
}

// 上传成功回调附件内容
function handleFinishCall(result: Attachment, success: boolean) {
  if (success) {
    reloadTable();
  }
}
</script>
```

### 开关 Switch
```vue
<template>
  <n-switch v-model:value="value" />
</template>

<script lang="ts" setup>
import { ref } from 'vue'
const value = ref(null);
</script>
```

### 评分 Rate
```vue
<template>
  <n-rate allow-half :default-value="value" />
</template>

<script lang="ts" setup>
import { ref } from 'vue'
const value = ref(null);
</script>
```

### 省市区选择器 CitySelector
```vue
<template>
  <CitySelector v-model:value="value" />
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import CitySelector from '@/components/CitySelector/citySelector.vue';
const value = ref(null);
</script>
```

### 图标选择器 IconSelector
```vue
<template>
  <IconSelector style="width: 100%" v-model:value="value" />
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import IconSelector from '@/components/IconSelector/index.vue';
const value = ref(null);
</script>
```


更多组件请参考：https://www.naiveui.com/zh-CN/os-theme/components/button


