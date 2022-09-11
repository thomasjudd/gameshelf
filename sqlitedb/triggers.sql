CREATE TRIGGER game_ai AFTER INSERT ON game BEGIN
  INSERT INTO game_fts(game_id, name, location) VALUES (new.game_id, new.name, new.location);
END;
CREATE TRIGGER game_ad AFTER DELETE ON game BEGIN
  INSERT INTO game_fts(game_fts, game_id, name, location) VALUES('delete', old.game_id, old.name, old.location);
END;
CREATE TRIGGER game_au AFTER UPDATE ON game BEGIN
  INSERT INTO game_fts(game_fts, game_id, name, location) VALUES('delete', old.game_id, old.name, old.location);
  INSERT INTO game_fts(game_id, name, location) VALUES (new.game_id, new.name, new.location);
END;
