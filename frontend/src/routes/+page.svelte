<script lang="ts">
  import "../global.css"
  import Hcaptcha from "$lib/components/Hcaptcha.svelte";
  //import { IconMoodHappyFilled, IconBrandTelegram } from "@tabler/icons-svelte"
	import { onMount } from "svelte";
  import { MessageType, type ChatMessage } from "$lib/type"

  let socket = $state<WebSocket>()
  let textareaRef: HTMLTextAreaElement | undefined
  let messages = $state<string[]>([])

  function onSuccess(resp: string){
    
  }

  function sendMessage(){
    if (socket && textareaRef && textareaRef.value !== "") {
      messages.push(`あなた: ${textareaRef.value}`)      
      socket.send(JSON.stringify({ content: textareaRef.value })) 

      textareaRef.value = ""
    }
  }

  onMount(() => {
    socket = new WebSocket("ws://localhost:3000/ws")

    socket.onopen = () => {
      console.log("Websocket connected")
    }

    socket.onmessage = (ev: MessageEvent<string>) => {
      const data = JSON.parse(ev.data) as ChatMessage
      messages.push(`${data.username}: ${data.content}`)
    }

    socket.onerror = (ev) => {
      console.error(ev)
    }

    socket.onclose = (ev) => {
      console.warn(`${ev.code}: ${ev.reason}`)
    }
  })
</script>

<main class="flex flex-col h-screen">
  <!-- <Hcaptcha onSuccess={onSuccess} /> -->
  <div class="flex flex-1 justify-center h-full">
    <div class="w-[65%] h-full overflow-y-auto p-4">
      {#if messages.length == 0}
        <h1>メッセージ無し</h1>
      {/if}
      {#each messages as msg}
        <div class="px-4 py-2 w-full">{msg}</div>
      {/each}
    </div>
  </div>
  <div class="flex justify-center items-center w-full py-2">
    <label for="chat" class="sr-only">Your message</label>
    <div class="flex justify-between items-center w-[60%] px-3 py-2 rounded-lg bg-gray-50">
      <button class="px-0 mr-4">😊</button>
      <textarea bind:this={textareaRef} id="chat" class="flex-1 resize-none outline-none " rows="1"></textarea>
      <button onclick={sendMessage} class="bg-green-500 active:scale-95">
        ✅
        <span>送信</span>
      </button>
    </div>
  </div>
</main>