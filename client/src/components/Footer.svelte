<script lang="ts">
  import LinkWithPhoto from "./LinkWithPhoto.svelte";
</script>

<div class="fixed bottom-0 border-t-[1px] border-t-[#F6F6F6] w-full h-[45px]">
  <div
    class="absolute bottom-[12px] px-[24px] text-center w-full flex justify-between text-[13px]"
  >
    <span class="text-[#888888] font-black">
      Made with <img alt="Smile" src="/images/smile.png" class="w-[12px] h-[12px] inline mb-[4px] mx-[2px]" /> By GDSC-ASU (JO)
    </span>

    <div>
      <LinkWithPhoto
        link="https://gdsc.community.dev/applied-science-private-university/"
        imagePath="/images/google.png"
      />
      <LinkWithPhoto
        link="https://www.linkedin.com/company/GDSCASU"
        imagePath="/images/linkedin.png"
      />
      <LinkWithPhoto
        link="https://www.facebook.com/DSCASU"
        imagePath="/images/facebook.png"
      />
      <LinkWithPhoto
        link="https://github.com/mbaraa/dsc_logo_generator"
        imagePath="/images/github.png"
      />
    </div>
  </div>
</div>
