<script lang="ts">
	import { onMount } from "svelte";

  interface Props {
    onSuccess?: (resp: string) => void,
    onOpen?: () => void,
    onExpired?: () => void,
    onError?: (err: HCaptchaErrorCallback) => void
    onClose?: () => void
  }

  let { 
    onSuccess,
    onOpen,
    onExpired,
    onError,
    onClose
  }: Props = $props()

  onMount(async() => {
    hcaptcha.render("h-captcha", { 
      sitekey: "2566ef2a-608e-403c-9a2a-e7268db25152",
      theme: "light",
      size: "compact",
      "callback": onSuccess,
      "open-callback": onOpen,
      "expired-callback": onExpired,
      "chalexpired-callback": onExpired,
      "error-callback": onError,
      "close-callback": onClose
    })
  })
</script>

<div class="fixed inset-0 flex items-center justify-center backdrop-blur-xs z-999">
  <div id="h-captcha" class="w-full"></div>
</div>