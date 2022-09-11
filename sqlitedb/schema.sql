CREATE TABLE game(
game_id integer primary key,
name, location);
CREATE VIRTUAL TABLE game_fts using fts5(game_id, name, location)
/* game_fts(game_id,name,location) */;
CREATE TABLE IF NOT EXISTS 'game_fts_data'(id INTEGER PRIMARY KEY, block BLOB);
CREATE TABLE IF NOT EXISTS 'game_fts_idx'(segid, term, pgno, PRIMARY KEY(segid, term)) WITHOUT ROWID;
CREATE TABLE IF NOT EXISTS 'game_fts_content'(id INTEGER PRIMARY KEY, c0, c1, c2);
CREATE TABLE IF NOT EXISTS 'game_fts_docsize'(id INTEGER PRIMARY KEY, sz BLOB);
CREATE TABLE IF NOT EXISTS 'game_fts_config'(k PRIMARY KEY, v) WITHOUT ROWID;
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
