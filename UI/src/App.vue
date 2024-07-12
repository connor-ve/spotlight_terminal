<script setup>
import { ref } from 'vue';
import axios from 'axios';  

const command = ref('');

function handleCommand() {
  if (command.value.trim()) {
    sendCommandToAPI(command.value);
    command.value = ''; 
  }
}

async function sendCommandToAPI(cmd) {
  const apiURL = 'https://your-api-endpoint.com/command'; 
  try {
    const response = await axios.post(apiURL, { command: cmd });
    console.log('API Response:', response.data);
  } catch (error) {
    console.error('API Request Failed:', error);
  }
}
import { onMounted, onUnmounted } from 'vue';

onMounted(() => {
  const handleKeyDown = (event) => {
    if (event.key === 'Escape') {
      window.ipcRenderer.closeCurrentWindow();
    }
  };

  window.addEventListener('keydown', handleKeyDown);
});

onUnmounted(() => {
  window.removeEventListener('keydown', handleKeyDown);
});
</script>

<template>
  <div id="app">
    <input type="text" name="command-bar" id="command-bar" placeholder="Type command...">
  </div>
</template>

<style scoped>
#app {
  width: 100%;
  height: 100%;
  margin: 0;
  padding: 0;
  display: flex;
  justify-content: center;
  align-items: center;
  background: #333;
}

#command-bar {
  padding: 10px;
  font-size: 20px;
  color: #fff;
  background: #222;
  border: none;
  outline: none;
  box-sizing: border-box;
}



body {
  width: 100%;
  height: 100%;
  overflow: hidden;
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  border: red 1px solid;
}

* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;

}
</style>
