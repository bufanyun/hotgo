<template>
  <div>
    <a-list
      itemLayout="horizontal"
      :dataSource="data"
    >
      <a-list-item slot="renderItem" slot-scope="item, index" :key="index">
        <a-list-item-meta>
          <a slot="title">{{ item.title }}</a>
          <span slot="description">
            <span class="security-list-description">{{ item.description }}</span>
            <span v-if="item.value"> : </span>
            <span class="security-list-value">{{ item.value }}</span>
          </span>
        </a-list-item-meta>
        <template v-if="item.actions">
          <a slot="actions" @click="item.actions.callback">{{ item.actions.title }}</a>
        </template>

      </a-list-item>
    </a-list>
    <!-- 修改密码抽屉 -->
    <update-password
      ref="updatePassword"
    />
  </div>
</template>

<script>
import UpdatePassword from './UpdatePassword'

export default {
  components: {
    UpdatePassword
  },
  data () {
    return {
      data: [
        { title: '账户密码', description: '当前密码强度', value: '强', actions: { title: '修改', callback: () => { this.$refs.updatePassword.handleUpdatePwd() } } }
      ]
    }
  }
}
</script>

<style scoped>

</style>
