<script lang="ts">
    import Center from "$lib/components/Center.svelte";
    import Footer from "$lib/components/Footer.svelte";
    import Header from "$lib/components/Header.svelte";

    export const prerender = true;
</script>

<div>
    <Header />
    <Center />
    <Footer />
</div>
