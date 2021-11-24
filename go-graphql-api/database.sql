CREATE TABLE Request (
  request_id SERIAL PRIMARY KEY,
  user_email VARCHAR (100) NOT NULL,
  request_time TIMESTAMP WITH TIME ZONE NOT NULL,
  request_type VARCHAR(50) NOT NULL CHECK (request_type IN ('error report', 'contact lister'))
);

INSERT INTO Request VALUES
(DEFAULT, 'tkkim@ucsd.edu', current_timestamp, 'error report');