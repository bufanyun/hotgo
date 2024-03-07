import { onMounted } from 'vue';

export default function userCreateScript(src: string) {
  const createScriptPromise = new Promise((resolve, reject) => {
    onMounted(() => {
      const script = document.createElement('script');
      script.type = 'text/javascript';
      script.onload = () => {
        resolve('');
      };
      script.onerror = (error) => {
        reject(error);
      };
      script.src = src;
      document.head.appendChild(script);
    });
  });
  return {
    createScriptPromise,
  };
}
