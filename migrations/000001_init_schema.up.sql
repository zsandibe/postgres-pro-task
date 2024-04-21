CREATE TABLE IF NOT EXISTS command(
    id SERIAL PRIMARY KEY,
    script VARCHAR(255) NOT NULL    
);


CREATE TABLE IF NOT EXISTS command_output(
    id SERIAL PRIMARY KEY,
    command_id INT NOT NULL,
    output VARCHAR(255),
    time TIMESTAMP NOT NULL,
    FOREIGN KEY (command_id) REFERENCES command (id)
);