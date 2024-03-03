<script lang="ts">
	import { applyAction, enhance } from '$app/forms';
	import { page } from '$app/stores';
	import { beans } from '$lib/stores';

	let show = false;

	let createBean = () => {
		return async ({ result }) => {
			console.log(result);
			// needed to update the form store
			await applyAction(result);
			const newBean = $page.form.bean;

			$beans = [...$beans, newBean];
			console.log($beans);
			show = false;
		};
	};
</script>

<!-- Id string `json:"id"` -->
<!-- Name     string `json:"name"` -->
<!-- Roaster  string `json:"roaster"` -->
<!-- Country  string `json:"country"` -->
<!-- Varietal string `json:"varietal"` -->
<!-- Process  string `json:"process"` -->
<!-- Altitude string `json:"altitude"` -->
<!-- Notes    string `json:"notes"` -->
<!-- Weight   float32 `json:"weight"` -->

{#if show}
	<div class="card">
		<form method="POST" action="?/create" use:enhance={createBean}>
			<div class="flex flex-row justify-between space-x-1">
				<label class="label">Name: <input class="input" name="name" /> </label>
				<label class="label">Roaster: <input class="input" name="roaster" /> </label>
			</div>
			<label class="label">Weight: <input class="input" name="weight" /> </label>
			<label class="label">Country: <input class="input" name="country" /> </label>
			<label class="label">Varietal <input class="input" name="varietal" /> </label>
			<label class="label">Altitude <input class="input" name="altitude" /> </label>
			<label class="label">Notes: <input class="input" name="notes" /> </label>
			<div class="flex flex-row justify-between">
				<button class="button" type="submit">Create</button>
				<button class="button" on:click={() => (show = false)}>Cancel</button>
			</div>
		</form>
	</div>
{/if}
<button on:click={() => (show = true)}>+ Add Bean</button>
