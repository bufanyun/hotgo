import { PageHeaderWrapper } from './components'

const PageView = {
  name: 'PageView',
  render () {
    return (
      <PageHeaderWrapper>
        <router-view />
      </PageHeaderWrapper>
    )
  }
}

export default PageView
