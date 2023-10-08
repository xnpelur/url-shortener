function submitForm() {
    var urlValue = document.getElementById("url").value;
    var data = { "url": urlValue };

    fetch("http://127.0.0.1/link", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(data)
    })
    .then(async response => {
        if (response.ok) {
            fetchData()
        } else {
            const errorResponse = await response.json();
            console.error("Form data submission failed. Error message: ", errorResponse.message);
        }
    })
    .catch(error => {
        console.error("An error occurred:", error);
    });
}

function fetchData() {
    fetch("http://127.0.0.1/link", {
        method: "GET"
    })
    .then(async response => {
        if (response.ok) {
            data = await response.json()
            if (data == null) {
                renderLinks([])
            } else {
                lastLinks = data.reverse().slice(0, 10)
                renderLinks(lastLinks)
            }
        } else {
            const errorResponse = await response.json();
            console.error("Something went wrong. Error message: ", errorResponse.message);
        }
    })
    .catch(error => {
        console.error("An error occurred:", error);
    });
}

function renderLinks(links) {
    const table = document.getElementById("url-table")
    const tableBody = table?.querySelector('tbody')

    while (tableBody.firstChild) {
        tableBody.removeChild(tableBody.firstChild);
    }
    
    links?.forEach(link => {
        let newRow = createNewRow("127.0.0.1/" + link.short, link.url)
        tableBody.appendChild(newRow)
    });
}

function createNewRow(firstUrl, secondUrl) {
    const newRow = document.createElement('tr');

    const td1 = document.createElement('td');
    const td2 = document.createElement('td');

    const a1 = document.createElement('a')
    const a2 = document.createElement('a')

    a1.textContent = firstUrl
    a1.href = "http://" + firstUrl
    
    a2.textContent = secondUrl
    a2.href = "http://" + secondUrl

    td1.appendChild(a1)
    td2.appendChild(a2)

    newRow.appendChild(td1);
    newRow.appendChild(td2);
    
    return newRow
}

fetchData()