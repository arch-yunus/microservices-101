-- microservices-101: Veritabanı İlklendirme Betiği
-- Bu SQL, konteyner baslatıldıgında otomatik olarak tablo yapılarını olusturur.

-- Product Tablosu
CREATE TABLE IF NOT EXISTS products (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Order Tablosu
CREATE TABLE IF NOT EXISTS orders (
    id UUID PRIMARY KEY,
    product_id UUID NOT NULL,
    quantity INTEGER NOT NULL,
    total DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

-- Örnek Veriler (Simülasyon için)
INSERT INTO products (id, name, description, price, created_at)
VALUES 
    ('550e8400-e29b-41d4-a716-446655440000', 'Go Gopher Mascot', 'En sevimli maskot', 25.99, NOW()),
    ('660e8400-e29b-41d4-a716-446655440001', 'Microservices Book', 'Elite handbook', 49.50, NOW());
