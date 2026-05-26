<template>
  <div class="main-container">
    <n-card title="打印图片" :header-style="{ padding: '5px' }" :content-style="{ padding: '0px' }">
      <template #header-extra>
        <n-button type="primary" size="small" @click="printImage">打印</n-button>
      </template>
      <div class="image-wrapper">
        <img :src="imagePath" />
      </div>
    </n-card>
    <n-card
      title="打印HTML"
      :header-style="{ padding: '5px' }"
      :content-style="{ padding: '0px' }"
      class="mt-4"
    >
      <template #header-extra>
        <n-button type="primary" size="small" @click="printHtml">打印</n-button>
      </template>
      <div id="htmlWrapper" class="flex justify-center html-wrapper align-center flex-direction">
        <n-table :data="dataList">
          <thead>
            <tr>
              <th>姓名</th>
              <th>年龄</th>
              <th>性别</th>
              <th>职业</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(item, index) of dataList" :key="index">
              <td>{{ item.name }}</td>
              <td>{{ item.age }}</td>
              <td>{{ item.gender }}</td>
              <td>{{ item.career }}</td>
            </tr>
          </tbody>
        </n-table>
      </div>
    </n-card>
  </div>
</template>

<script lang="ts">
  import printJS from 'print-js';
  import imagePath from '@/assets/images/logo.png';
  import { defineComponent } from 'vue';
  export default defineComponent({
    name: 'Print',
    setup() {
      function printImage() {
        printJS({
          printable: imagePath,
          type: 'image',
          showModal: false,
        });
      }
      function printHtml() {
        printJS({
          printable: 'htmlWrapper',
          type: 'html',
          targetStyles: ['*'],
        });
      }
      return {
        printImage,
        printHtml,
        imagePath,
        dataList: [
          {
            name: '张三',
            age: 30,
            gender: '男',
            career: '工程师',
          },
          {
            name: '李四',
            age: 20,
            gender: '男',
            career: '服务员',
          },
          {
            name: '王五',
            age: 40,
            gender: '女',
            career: '售货员',
          },
        ],
      };
    },
  });
</script>

<style lang="less" scoped>
  .image-wrapper {
    width: 30%;
    margin: 0 auto;
    & > img {
      width: 100%;
    }
  }
  .html-wrapper {
    width: 80%;
    margin: 0 auto;
    & > h1 {
      color: red;
    }
  }
</style>
