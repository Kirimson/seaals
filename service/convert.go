package service

import (
	"database/sql"
	"encoding/json"
	"errors"
	"seaals/db"
	"seaals/models"
)

func convertSeal(s db.Seal, err error) (*models.Seal, error) {
	// Return early if main DB get failed
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var seal models.Seal
	sj, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(sj, &seal)
	return &seal, nil
}

func convertSeals(ss []db.Seal, err error) ([]*models.Seal, error) {
	var seals []*models.Seal
	for _, s := range ss {
		sm, err := convertSeal(s, err)
		if err != nil {
			return nil, err
		}
		seals = append(seals, sm)
	}
	return seals, nil
}

func convertTag(t db.Tag, err error) (*models.Tag, error) {
	// Return early if main DB get failed
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var tag models.Tag
	tj, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(tj, &tag)
	return &tag, nil
}

func convertTags(tt []db.Tag, err error) ([]*models.Tag, error) {
	var tags []*models.Tag
	for _, t := range tt {
		tm, err := convertTag(t, err)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tm)
	}
	return tags, nil
}

func convertTagStat(ts db.GetPopularTagsRow, err error) (*models.TagStat, error) {
	// Return early if main DB get failed
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var tagStat models.TagStat
	tsj, err := json.Marshal(ts)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(tsj, &tagStat)
	return &tagStat, nil
}

func convertTagStats(tss []db.GetPopularTagsRow, err error) ([]*models.TagStat, error) {
	var tagStats []*models.TagStat
	for _, t := range tss {
		tm, err := convertTagStat(t, err)
		if err != nil {
			return nil, err
		}
		tagStats = append(tagStats, tm)
	}
	return tagStats, nil
}
