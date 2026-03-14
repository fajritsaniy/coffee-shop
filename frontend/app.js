const API_BASE = 'http://localhost:3001/api/v1';
const API_KEY = 'RAHASIA';

let currentCategory = null;
let cart = [];

// DOM Elements
const categoriesList = document.getElementById('categories-list');
const menuGrid = document.getElementById('menu-grid');
const cartSidebar = document.getElementById('cart-sidebar');
const cartItemsContainer = document.getElementById('cart-items');
const cartTotalElement = document.getElementById('cart-total');
const cartToggle = document.getElementById('cart-toggle');
const closeCart = document.getElementById('close-cart');
const checkoutBtn = document.getElementById('checkout-btn');
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
        renderCategories(response.data);
    }
}

async function fetchMenuItems(categoryId = null) {
    const endpoint = categoryId
        ? `/menu-items-by-category/${categoryId}`
        : '/menu-items';

    const response = await apiFetch(endpoint);
    if (response && response.data) {
        renderMenu(response.data);
    }
}

function renderCategories(categories) {
    categoriesList.innerHTML = `
        <button class="category-btn ${!currentCategory ? 'active' : ''}" data-id="all">ALL</button>
    ` + categories.map(cat => `
        <button class="category-btn ${currentCategory === cat.id ? 'active' : ''}" data-id="${cat.id}">
            ${cat.name.toUpperCase()}
        </button>
    `).join('');

    document.querySelectorAll('.category-btn').forEach(btn => {
        btn.addEventListener('click', (e) => {
            const id = e.target.dataset.id;
            currentCategory = id === 'all' ? null : parseInt(id);

            document.querySelectorAll('.category-btn').forEach(b => b.classList.remove('active'));
            e.target.classList.add('active');

            fetchMenuItems(currentCategory);
        });
    });
}

function renderMenu(items) {
    menuGrid.innerHTML = items.map(item => `
        <div class="menu-item">
            ${item.image_url ? `<img src="${item.image_url}" alt="${item.name}" class="item-image">` : '<div class="item-image"></div>'}
            <div class="item-content">
                <div>
                    <h3>${item.name.toUpperCase()}</h3>
                    <p>${item.description || 'No description available.'}</p>
                </div>
                <div class="item-footer">
                    <span class="price">$${(item.price / 100).toFixed(2)}</span>
                    <button onclick="addToCart(${item.id}, '${item.name}', ${item.price})">
                        <i data-lucide="shopping-cart"></i> ADD
                    </button>
                </div>
            </div>
        </div>
    `).join('');
    lucide.createIcons();
}

function addToCart(id, name, price) {
    const existing = cart.find(item => item.id === id);
    if (existing) {
        existing.quantity += 1;
    } else {
        cart.push({ id, name, price, quantity: 1 });
    }
    updateCartUI();
    showNotification(`Added ${name} to cart`);
}

function updateCartUI() {
    cartToggle.innerText = `CART (${cart.reduce((sum, item) => sum + item.quantity, 0)})`;

    cartItemsContainer.innerHTML = cart.map((item, index) => `
        <div class="cart-item">
            <div>
                <strong>${item.name.toUpperCase()}</strong>
                <div>x${item.quantity} - $${((item.price * item.quantity) / 100).toFixed(2)}</div>
            </div>
            <button onclick="removeFromCart(${index})">REMOVE</button>
        </div>
    `).join('');

    const total = cart.reduce((sum, item) => sum + (item.price * item.quantity), 0);
    cartTotalElement.innerText = `$${(total / 100).toFixed(2)}`;
}

function removeFromCart(index) {
    cart.splice(index, 1);
    updateCartUI();
}

function setupEventListeners() {
    cartToggle.addEventListener('click', () => cartSidebar.classList.remove('hidden'));
    closeCart.addEventListener('click', () => cartSidebar.classList.add('hidden'));

    checkoutBtn.addEventListener('click', async () => {
        if (cart.length === 0) return;

        const orderData = {
            table_id: 1, // Default table for demo
            items: cart.map(item => ({
                menu_item_id: item.id,
                quantity: item.quantity
            }))
        };

        const response = await apiFetch('/orders', {
            method: 'POST',
            body: JSON.stringify(orderData)
        });

        if (response) {
            showNotification('Order placed successfully!');
            cart = [];
            updateCartUI();
            cartSidebar.classList.add('hidden');
        }
    });
}

function showNotification(message) {
    notification.innerText = message;
    notification.classList.remove('hidden');
    setTimeout(() => {
        notification.classList.add('hidden');
    }, 3000);
}

init();
