<template>
  <component :is="component" :class="className" aria-hidden="true">
    <use :href="iconName" />
  </component>
</template>

<script lang="ts">
  import { defineComponent } from 'vue';

  export default defineComponent({
    name: 'SvgIcon',
    props: {
      prefix: {
        type: String,
        default: 'icon',
      },
      name: {
        type: String,
        required: true,
      },
    },
    computed: {
      component(): string {
        return this.prefix === 'icon' ? 'svg' : 'i';
      },
      iconName(): string {
        return `#${this.prefix}-${this.name}`;
      },
      className(): string {
        if (this.prefix === 'icon') {
          return 'svg-icon';
        } else if (this.prefix === 'iconfont') {
          return 'iconfont icon-' + this.name;
        } else {
          return '';
        }
      },
    },
  });
</script>

<style scoped>
  .svg-icon {
    width: 1em;
    height: 1em;
    vertical-align: -0.15em;
    fill: currentColor;
    overflow: hidden;
  }
  .svg-icon:hover {
    fill: var(--primary-color-hover);
  }

  .svg-external-icon {
    background-color: currentColor;
    mask-size: cover !important;
    display: inline-block;
  }
</style>
