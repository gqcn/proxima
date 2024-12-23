// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"proxima/app/user/internal/dao/internal"
)

// internalWordsDao is an internal type for wrapping the internal DAO implementation.
type internalWordsDao = *internal.WordsDao

// wordsDao is the data access object for the table words.
// You can define custom methods on it to extend its functionality as needed.
type wordsDao struct {
	internalWordsDao
}

var (
	// Words is a globally accessible object for table words operations.
	Words = wordsDao{
		internal.NewWordsDao(),
	}
)

// Add your custom methods and functionality below.