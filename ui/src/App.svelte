<script>
	let code = '';
	let url = '';
	$: listofurls = [];
	
	async function addURL () {
		const res = await fetch('/', {
			method: 'POST',
			body: JSON.stringify({
				"code": code,
				"url": url
			}),
			headers: {
				"content-type":"application/json"
			}
		})
	}
	
	async function getUrls(){
			const res = await fetch("/?list=true");
			let urls = await res.json();
			console.log(urls)
			listofurls = urls
	}
	getUrls()

	setInterval(getUrls, 2000)
</script>

<h2>Shorty</h2>

<input type="text" placeholder="code" bind:value={code}>
<input type="text" placeholder="url" bind:value={url}>
<button on:click={addURL}>Add URL</button>

<br><br>
<table>
	<thead>
		<th>Code</th>
		<th>URL</th>
	</thead>
	<tbody>
		{#each listofurls as item}
			<tr>
				<td>{item.code}</td>
				<td>{item.url}</td>
			</tr>
		{/each}
	</tbody>
</table>