let saysInput = document.getElementById("says-input");
let saysLink = document.getElementById("says-link");
// Reset input
saysInput.value = "Something";
UpdateSays();

getStats();

function UpdateSays() {
  saysLink.href = `/seal/says/${saysInput.value}`;
}

async function getStats() {
  const url = "/api/stats";
  let countText = document.getElementById("seal-count");
  let popularTags = document.getElementById("popular-tags");
  try {
    const response = await fetch(url);
    if (!response.ok) {
      throw new Error(`Response status: ${response.status}`);
    }

    const stats = await response.json();
    countText.innerHTML = stats["count"];
    stats.tags.forEach((tag, idx) => {
      row = popularTags.insertRow(idx);
      nameCell = row.insertCell(0);
      nameCell.innerHTML = tag.name;
      countCell = row.insertCell(1);
      countCell.innerHTML = tag.count;

      exampleLink = document.createElement("a");
      exampleLink.href = `/seal?tag=${tag.name}`;
      exampleLink.target = "_blank";
      exampleButton = document.createElement("button");
      exampleButton.innerHTML = "Get a Seal!";
      exampleButton.classList.add(
        "rounded-md",
        "py-2",
        "px-4",
        "border",
        "border-transparent",
        "transition-all",
        "shadow-md",
        "hover:shadow-lg",
        "text-center",
        "text-black",
        "dark:text-white",
        "focus:shadow-none",
        "bg-sky-500",
        "dark:bg-sky-800",
        "hover:bg-sky-300",
        "hover:dark:bg-sky-700",
        "hover:cursor-pointer",
      );
      exampleCell = row.insertCell(2);
      exampleLink.appendChild(exampleButton);
      exampleCell.appendChild(exampleLink);
    });

    // Use the most popular tag as the tag example
    tagExampleLink = document.getElementById("tag-example");
    tagExampleLink.href = `/seal?tag=${stats.tags[0].name}`;
    tagExampleBtn = document.getElementById("tag-example-btn");
    tagExampleBtn.innerHTML = stats.tags[0].name;
  } catch (error) {
    console.error(error.message);
  }
}
