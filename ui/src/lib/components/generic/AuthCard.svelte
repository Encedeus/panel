<script lang="ts">
  export let height: string;
  export let className = "";
</script>

<div class="drop-shadow-xl {className}">
  <div class="bg-indigo-900 h-20 w-[35rem] rounded-t-xl flex justify-start items-center">
    <span class="ml-8">
      <slot name="title"/>
    </span>
  </div>
  <div class="flex flex-col gap-8 items-center justify-end bg-indigo-950 rounded-b-xl h-{height} w-[35rem]">
    <div class="mt-7">
      <slot name="inputs"/>
    </div>
    <div class="mb-11">
      <slot name="button"/>
    </div>
  </div>
</div>
