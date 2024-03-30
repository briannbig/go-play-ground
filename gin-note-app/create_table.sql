DROP DATABASE IF EXISTS gin_note_app;
CREATE DATABASE gin_note_app;
USE gin_note_app;
DROP TABLE IF EXISTS notes;
CREATE TABLE notes(
    id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    title VARCHAR(255),
    content TEXT,
    time_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
INSERT INTO notes (title, content)
VALUES (
        'Meeting Agenda',
        'Agenda for the weekly team meeting: 1. Project updates, 2. Action items, 3. Discussion on upcoming deadlines.'
    ),
    (
        'Shopping List',
        'Grocery shopping list: 1. Milk, 2. Bread, 3. Eggs, 4. Vegetables.'
    ),
    (
        'Travel Itinerary',
        'Itinerary for upcoming trip: 1. Flight details, 2. Hotel reservation, 3. Sightseeing plans.'
    );