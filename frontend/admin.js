const API_BASE = 'http://localhost:3001/api/v1';
const API_KEY = 'RAHASIA';

let categories = [];

// DOM Elements
const menuList = document.getElementById('menu-list');
const addItemBtn = document.getElementById('add-item-btn');
const itemModal = document.getElementById('item-modal');
const itemForm = document.getElementById('item-form');
const closeModal = document.getElementById('close-modal');
const categorySelect = document.getElementById('item-category');
const notification = document.getElementById('notification');

// Initialize
async function init() {
    await fetchCategories();
    await fetchMenuItems();
    setupEventListeners();
    lucide.createIcons();
}

async function apiFetch(endpoint, options = {}) {
    const url = `${API_BASE}${endpoint}`;
    const defaultOptions = {
        headers: {
            'X-Api-Key': API_KEY,
            'Content-Type': 'application/json'
        }
    };

    try {
        const response = await fetch(url, { ...defaultOptions, ...options });
        if (!response.ok) throw new Error(`API Error: ${response.status}`);
        return await response.json();
    } catch (error) {
        showNotification(`Error: ${error.message}`);
        console.error(error);
        return null;
    }
}

async function fetchCategories() {
    const response = await apiFetch('/menu-categories');
    if (response && response.data) {
        categories = response.data;
        categorySelect.innerHTML = categories.map(cat => `
            <option value="${cat.id}">${cat.name}</option>
        `).join('');
    }
}

async function fetchMenuItems() {
    const response = await apiFetch('/menu-items');
    if (response && response.data) {
        renderTable(response.data);
    }
}

function renderTable(items) {
    menuList.innerHTML = items.map(item => `
        <tr>
            <td>
                <img src="${item.image_url || 'https://via.placeholder.com/50'}" class="admin-thumb" alt="${item.name}">
            </td>
            <td><strong>${item.name}</strong></td>
            <td>${getCategoryName(item.category_id)}</td>
            <td>$${(item.price / 100).toFixed(2)}</td>
            <td>
                <span class="status-badge ${item.is_available ? 'available' : 'unavailable'}">
                    ${item.is_available ? 'In Stock' : 'Out of Stock'}
                </span>
            </td>
            <td>
                <div class="actions">
                    <button class="action-btn edit" onclick="editItem(${JSON.stringify(item).replace(/"/g, '&quot;')})">
                        <i data-lucide="edit-2"></i>
                    </button>
                    <button class="action-btn delete" onclick="deleteItem(${item.id})">
                        <i data-lucide="trash-2"></i>
                    </button>
                </div>
            </td>
        </tr>
    `).join('');
    lucide.createIcons();
}

function getCategoryName(id) {
    const cat = categories.find(c => c.id === id);
    return cat ? cat.name : 'Unknown';
}

function setupEventListeners() {
    addItemBtn.addEventListener('click', () => {
        itemForm.reset();
        document.getElementById('item-id').value = '';
        document.getElementById('modal-title').innerText = 'Add Menu Item';
        itemModal.classList.remove('hidden');
    });

    closeModal.addEventListener('click', () => {
        itemModal.classList.add('hidden');
    });

    itemForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        const id = document.getElementById('item-id').value;
        const data = {
            category_id: parseInt(document.getElementById('item-category').value),
            name: document.getElementById('item-name').value,
            price: parseInt(document.getElementById('item-price').value),
            description: document.getElementById('item-description').value,
            image_url: document.getElementById('item-image-url').value,
            is_available: true
        };

        const method = id ? 'PUT' : 'POST';
        const url = id ? `/menu-items/${id}` : '/menu-items';

        const response = await apiFetch(url, {
            method: method,
            body: JSON.stringify(data)
        });

        if (response) {
            showNotification(id ? 'Item updated!' : 'Item added!');
            itemModal.classList.add('hidden');
            fetchMenuItems();
        }
    });
}

window.editItem = (item) => {
    document.getElementById('item-id').value = item.id;
    document.getElementById('item-name').value = item.name;
    document.getElementById('item-category').value = item.category_id;
    document.getElementById('item-price').value = item.price;
    document.getElementById('item-description').value = item.description;
    document.getElementById('item-image-url').value = item.image_url;
    document.getElementById('modal-title').innerText = 'Edit Menu Item';
    itemModal.classList.remove('hidden');
};

window.deleteItem = async (id) => {
    if (!confirm('Are you sure you want to delete this item?')) return;

    const response = await apiFetch(`/menu-items/${id}`, { method: 'DELETE' });
    if (response) {
        showNotification('Item deleted');
        fetchMenuItems();
    }
};

function showNotification(message) {
    notification.innerText = message;
    notification.classList.remove('hidden');
    setTimeout(() => notification.classList.add('hidden'), 3000);
}

init();
