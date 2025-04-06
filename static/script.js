// Fetch artist data from the Groupie Tracker API
fetch("/api/artists")  // Assuming you're using a local proxy through Go
  .then(res => res.json())           // Convert the response to JSON
  .then(artists => {
    const container = document.getElementById("artist-list");
    container.innerHTML = "";        // Clear "Loading artists..."

    artists.forEach(artist => {
      const card = document.createElement("div");
      card.style.border = "1px solid #ccc";
      card.style.margin = "10px";
      card.style.padding = "10px";
      card.style.borderRadius = "8px";

      card.innerHTML = `
        <h2>${artist.name}</h2>
        <img src="${artist.image}" alt="${artist.name}" width="200">
        <p><strong>Members:</strong> ${artist.members.join(", ")}</p>
        <p><strong>Created:</strong> ${artist.creationDate}</p>
        <p><strong>First Album:</strong> ${artist.firstAlbum}</p>
      `;

      container.appendChild(card);
    });
  })
  .catch(err => {
    console.error("Error fetching artists:", err);
    document.getElementById("artist-list").innerText = "Failed to load artist data.";
  });
