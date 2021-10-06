<script>
	let code = "";
	let url = "";
	$: listofurls = [];
	let taglines = [
		"Zuckerberg is Sus",
		"12345 is a bad password!",
		"90% bug free!",
		"Any computer is a laptop if you're brave enough!",
		"Child's play!",
		"Cogito ergo sum!",
		"Supercalifragilisticexpialidocious!",
	];
	$: tagline = taglines[Math.floor(Math.random() * taglines.length)];
	$: href = "";

	async function addURL() {
		const res = await fetch("/", {
			method: "POST",
			body: JSON.stringify({
				code: code,
				url: url,
			}),
			headers: {
				"content-type": "application/json",
			},
		});
	}

	async function deleteURL(code) {
		const res = await fetch("/", {
			method: "DELETE",
			body: JSON.stringify({
				code: code,
			}),
			headers: {
				"content-type": "application/json",
			},
		});
	}

	function copyURL(code) {
		document.getElementById("icon-" + code).classList.replace("far", "fas");
		navigator.clipboard.writeText("https://" + href + code);
		setTimeout(function () {
			document
				.getElementById("icon-" + code)
				.classList.replace("fas", "far");
		}, 1000);
	}

	async function getUrls() {
		const res = await fetch("/?list=true");
		let urls = await res.json();
		listofurls = urls;
	}

	async function getDomain() {
		const res = await fetch("/domain");
		let domain = await res.text();
		console.log(domain);
		href = domain;
	}
	getUrls();
	getDomain();

	setInterval(getUrls, 2000);
</script>

<div class="columns">
	<div class="column is-3" />
	<div class="column">
		<!-- PAGE TITLE -->
		<section class="hero is-info" style="background: #22a6b3;">
			<div class="hero-body">
				<div class="container has-text-centered">
					<p class="title">Shortie - The URL Shortener</p>
					<p class="subtitle">
						{tagline}
					</p>
				</div>
			</div>
			<div class="hero-foot">
				<div class="is-flex is-justify-content-center">
					<div class="field has-addons p-3">
						<p class="control">
							<a class="button is-static is-primary is-light">
								{href}
							</a>
						</p>
						<p class="control">
							<input
								class="input is-primary"
								type="text"
								placeholder="Short URL"
								bind:value={code}
							/>
						</p>
					</div>
					<div class="field has-addons p-3">
						<p class="control">
							<input
								class="input is-primary"
								type="text"
								placeholder="Full URL"
								bind:value={url}
							/>
						</p>
						<p class="control">
							<button
								class="button is-primary is-light"
								on:click={addURL}
								style="background: #7ed6df;"
							>
								Shortie!
							</button>
						</p>
					</div>
				</div>
			</div>
		</section>
		<br />
		{#each listofurls as item}
			<div class="box">
				<!-- <div class="card-content"> -->
				<div class="columns">
					<div class="column is-9">
						<p class="is-size-4">
							<a href={"https://" + href + item.code}
								>{href}{item.code}</a
							>
						</p>
						<p class="is-size-6">
							{item.url}
						</p>
					</div>
					<div class="column is-3">
						<div class="field has-addons has-text-centered">
							<div class="control">
								<button
									class="button is-primary is-inverted is-medium"
									id={item.code}
									on:click={copyURL(item.code)}
								>
									<i
										class="far fa-copy"
										id={"icon-" + item.code}
									/>
								</button>
							</div>
							<div class="control">
								<button
									class="button is-danger is-inverted is-medium"
									id={item.code}
									on:click|stopPropagation={deleteURL(
										item.code
									)}
								>
									<i class="far fa-trash-alt" /></button
								>
							</div>
						</div>
					</div>
				</div>
				<!-- </div> -->
			</div>
		{/each}
	</div>
	<div class="column is-3" />
</div>