<script lang="ts">
  export let link: string;
  export let imagePath: string;
</script>

<a href={link} target="_blank">
  <img alt={link} src={imagePath} class="w-[16px] h-[16px] inline mx-[4px]" />
</a>
