<script lang="ts">
  export let size: "sm" | "md" | "lg" = "md";
  export let className: string;

  const sizes = new Map<string, string>([
      ["sm"], ["w-64 h-3"],
      ["md", "w-72 h-3"],
      ["lg", "w-[40vw] h-10 rounded-full px-12"],
  ]);
</script>

<div>
  <input type="search" placeholder="Search" class="{sizes.get(size)} placeholder:text-xs placeholder-white placeholder-opacity-90 font-inter font-normal text-sm text-indigo-300 focus:outline-indigo-400 outline-none outline-offset-0 {className}">
</div>

<style>
  input {
      background: url("data:image/svg+xml, %3Csvg width='20' height='20' viewBox='0 0 16 16' fill='none' xmlns='http://www.w3.org/2000/svg'%3E%3Cpath d='M14 14L10.5354 10.5353M10.5354 10.5353C11.4731 9.59762 11.9999 8.3258 11.9999 6.99967C11.9999 5.67354 11.4731 4.40172 10.5354 3.464C9.59765 2.52629 8.32583 1.99948 6.9997 1.99948C5.67357 1.99948 4.40175 2.52629 3.46403 3.464C2.52632 4.40172 1.99951 5.67354 1.99951 6.99967C1.99951 8.3258 2.52632 9.59762 3.46403 10.5353C4.40175 11.4731 5.67357 11.9999 6.9997 11.9999C8.32583 11.9999 9.59765 11.4731 10.5354 10.5353Z' stroke='white' stroke-width='1.5' stroke-linecap='round' stroke-linejoin='round'/%3E%3C/svg%3E") no-repeat rgb(79 70 229);
      background-position-y: center;
      background-position-x: 1.25rem;
  }
</style>