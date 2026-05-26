<template>
  <div>
    <n-modal
      v-model:show="isShowModal"
      :show-icon="false"
      preset="dialog"
      :title="params?.id > 0 ? '编辑 #' + params?.id : '添加'"
      :style="{
        width: dialogWidth,
      }"
    >
      <n-form
        :model="params"
        :rules="rules"
        ref="formRef"
        label-placement="left"
        :label-width="80"
        class="py-4"
      >
        <n-form-item label="标题" path="title">
          <n-input placeholder="请输入标题" v-model:value="params.title" />
        </n-form-item>

        <n-form-item label="分类ID" path="categoryId">
          <n-input-number placeholder="请输入分类ID" v-model:value="params.categoryId" />
        </n-form-item>

        <n-form-item label="标签" path="flag">
          <n-checkbox-group v-model:value="params.flag">
            <n-space>
              <n-checkbox
                v-for="item in dict.getOptionUnRef('sys_notice_type')"
                :key="item.value"
                :value="item.value"
                :label="item.label"
              />
            </n-space>
          </n-checkbox-group>
        </n-form-item>

        <n-form-item label="描述" path="description">
          <n-input type="textarea" placeholder="描述" v-model:value="params.description" />
        </n-form-item>

        <n-form-item label="内容" path="content">
          <Editor style="height: 450px" v-model:value="params.content" />
        </n-form-item>

        <n-form-item label="单图" path="image">
          <UploadImage v-model:value="params.image" />
        </n-form-item>

        <n-form-item label="多图" path="images">
          <FileChooser :maxNumber="10" v-model:value="params.images" fileType="image" />
        </n-form-item>

        <n-form-item label="单附件" path="attachfile">
          <FileChooser v-model:value="params.attachfile" />
        </n-form-item>

        <n-form-item label="多附件" path="attachfiles">
          <UploadFile :maxNumber="10" v-model:value="params.attachfiles" />
        </n-form-item>

        <n-form-item label="键值对" path="map">
          <n-dynamic-input
            v-model:value="params.map"
            preset="pair"
            key-placeholder="键名"
            value-placeholder="键值"
          />
        </n-form-item>
        <n-form-item label="推荐星" path="star">
          <n-rate allow-half :default-value="params.star" :on-update:value="updateStar" />
        </n-form-item>

        <n-form-item label="价格" path="price">
          <n-input-number
            placeholder="请输入价格"
            clearable
            v-model:value="params.price"
            :precision="2"
          />
        </n-form-item>

        <n-form-item label="活动时间" path="activityAt">
          <DatePicker v-model:formValue="params.activityAt" type="date" />
        </n-form-item>

        <n-form-item label="开放时间" path="startAt">
          <DatePicker
            v-model:startValue="params.startAt"
            v-model:endValue="params.endAt"
            type="datetimerange"
          />
        </n-form-item>

        <n-form-item label="用户渠道" path="channel">
          <n-select
            v-model:value="params.channel"
            :options="dict.getOptionUnRef('sys_user_channel')"
          />
        </n-form-item>

        <n-form-item label="所在城市" path="cityId">
          <CitySelector v-model:value="params.cityId" />
        </n-form-item>

        <n-form-item label="用户爱好" path="hobby">
          <n-select
            multiple
            v-model:value="params.hobby"
            :options="dict.getOptionUnRef('sys_user_hobby')"
          />
        </n-form-item>

        <n-form-item label="QQ" path="qq">
          <n-input placeholder="请输入QQ号" v-model:value="params.qq" />
        </n-form-item>

        <n-form-item label="邮箱" path="email">
          <n-input placeholder="请输入邮箱地址" v-model:value="params.email" />
        </n-form-item>

        <n-form-item label="手机号" path="mobile">
          <n-input placeholder="请输入手机号" v-model:value="params.mobile" />
        </n-form-item>

        <n-form-item label="排序" path="sort">
          <n-input-number v-model:value="params.sort" clearable />
        </n-form-item>

        <n-form-item label="状态" path="status">
          <n-radio-group v-model:value="params.status" name="status">
            <n-radio-button
              v-for="status in dict.getOptionUnRef('sys_normal_disable')"
              :key="Number(status.value)"
              :value="Number(status.value)"
              :label="status.label"
            />
          </n-radio-group>
        </n-form-item>

        <n-form-item label="备注" path="remark">
          <n-input type="textarea" placeholder="请输入备注" v-model:value="params.remark" />
        </n-form-item>
      </n-form>

      <template #action>
        <n-space>
          <n-button @click="closeForm">取消</n-button>
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm">确定</n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
  import { ref, computed, watch } from 'vue';
  import { rules, State, newState } from './model';
  import { Edit, MaxSort } from '@/api/addons/hgexample/table';
  import { useMessage } from 'naive-ui';
  import { adaModalWidth } from '@/utils/hotgo';
  import DatePicker from '@/components/DatePicker/datePicker.vue';
  import Editor from '@/components/Editor/editor.vue';
  import UploadImage from '@/components/Upload/uploadImage.vue';
  import UploadFile from '@/components/Upload/uploadFile.vue';
  import CitySelector from '@/components/CitySelector/citySelector.vue';
  import FileChooser from '@/components/FileChooser/index.vue';
  import { useDictStore } from '@/store/modules/dict';
  const emit = defineEmits(['reloadTable', 'updateShowModal']);

  interface Props {
    showModal: boolean;
    formParams?: State;
  }

  const props = withDefaults(defineProps<Props>(), {
    showModal: false,
    formParams: () => {
      return newState(null);
    },
  });

  const isShowModal = computed({
    get: () => {
      return props.showModal;
    },
    set: (value) => {
      emit('updateShowModal', value);
    },
  });

  const params = computed(() => {
    return props.formParams;
  });

  const dict = useDictStore();
  const message = useMessage();
  const formRef = ref<any>({});
  const formBtnLoading = ref(false);
  const dialogWidth = computed(() => {
    return adaModalWidth();
  });

  function confirmForm(e) {
    e.preventDefault();
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        Edit(params.value).then((_res) => {
          message.success('操作成功');
          setTimeout(() => {
            isShowModal.value = false;
            emit('reloadTable');
          });
        });
      } else {
        message.error('请填写完整信息');
      }
      formBtnLoading.value = false;
    });
  }

  function updateStar(num) {
    params.value.star = num;
  }

  function closeForm() {
    isShowModal.value = false;
  }

  watch(
    () => params.value,
    (value) => {
      if (value.id === 0) {
        MaxSort().then((res) => {
          params.value.sort = res.sort;
        });
      }
    }
  );
</script>

<style lang="less"></style>
