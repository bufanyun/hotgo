/* eslint-disable */
class SideEffect {
  constructor ({ propsToState, handleStateChange }) {
    if (typeof propsToState !== 'function') {
      throw new Error('Expected propsToState to be a function.')
    }
    if (typeof handleStateChange !== 'function') {
      throw new Error('Expected handleStateChange to be a function.')
    }
    this.options = {
      propsToState,
      handleStateChange
    }
  }
}

export default SideEffect
