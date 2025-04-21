<script lang="ts">
	import { onMount } from 'svelte';
    import type { PageData } from './$types';
    let { data }: { data: PageData } = $props();
    let refresh = $state(3000);
    let count = $state(0);
    let respond: any = $state(data);
    const fetchData = () => {
        count++;
        fetch("/api/resources").then(async (res) => {
            respond = await res.json();
            console.log(respond);
        }).catch((err) => {
            console.log(err);
        });
        setTimeout(fetchData, refresh);
    }
    onMount(() => {
        fetchData();
    });
</script>


<h1>Dashboard</h1>
{#if data}
    CPU:
    {#each respond?.cpu as cpu, idx}
        <div>CPU #{idx}: {Math.round(cpu).toFixed(2)}</div>
    {/each}
    Memory:
    {#if respond?.mem}
        <div>
            <span>{Math.round(respond?.mem?.usedPercent).toFixed(2)} / 100%</span> used /<span>{Math.round(respond?.mem?.total / (1024 ** 3)).toFixed(2)} GB total</span>
        </div>
    {/if}
{/if}
{count}

<button onclick={() => fetchData()}>Refresh</button>

<input type="number" bind:value={refresh} min="1000" />

{refresh}