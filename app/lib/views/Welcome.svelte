<script lang="ts">
  import Layout from "$lib/components/Layout.svelte";
  import { action } from "$lib/scripts/core/action.ts";
  import type { Form as Props } from "$lib/types/gen/main/lib/session/memory/Form";

  let { Username, Error }: Props = $props();
</script>

<Layout title="Welcome">
  <form method="POST" {...action("/")}>
    <input type="text" value="{Username}" placeholder="Username" name="username" class="input input-neutral" />
    <div class="pt-6"></div>

    {#if Error}
        <div class="text-error">{Error}</div>
        <div class="pt-6"></div>
    {/if}
    <button class="btn btn-secondary">Submit</button>
  </form>
</Layout>
