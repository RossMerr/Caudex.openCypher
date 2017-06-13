// Generated from Cypher.g4 by ANTLR 4.7.

package openCypher // Cypher
import "github.com/antlr/antlr4/runtime/Go/antlr"

// CypherListener is a complete listener for a parse tree produced by CypherParser.
type CypherListener interface {
	antlr.ParseTreeListener

	// EnterCypher is called when entering the cypher production.
	EnterCypher(c *CypherContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterRegularQuery is called when entering the regularQuery production.
	EnterRegularQuery(c *RegularQueryContext)

	// EnterUnion is called when entering the union production.
	EnterUnion(c *UnionContext)

	// EnterSingleQuery is called when entering the singleQuery production.
	EnterSingleQuery(c *SingleQueryContext)

	// EnterSinglePartQuery is called when entering the singlePartQuery production.
	EnterSinglePartQuery(c *SinglePartQueryContext)

	// EnterReadOnlyEnd is called when entering the readOnlyEnd production.
	EnterReadOnlyEnd(c *ReadOnlyEndContext)

	// EnterReadUpdateEnd is called when entering the readUpdateEnd production.
	EnterReadUpdateEnd(c *ReadUpdateEndContext)

	// EnterUpdatingEnd is called when entering the updatingEnd production.
	EnterUpdatingEnd(c *UpdatingEndContext)

	// EnterMultiPartQuery is called when entering the multiPartQuery production.
	EnterMultiPartQuery(c *MultiPartQueryContext)

	// EnterReadPart is called when entering the readPart production.
	EnterReadPart(c *ReadPartContext)

	// EnterUpdatingPart is called when entering the updatingPart production.
	EnterUpdatingPart(c *UpdatingPartContext)

	// EnterUpdatingStartClause is called when entering the updatingStartClause production.
	EnterUpdatingStartClause(c *UpdatingStartClauseContext)

	// EnterUpdatingClause is called when entering the updatingClause production.
	EnterUpdatingClause(c *UpdatingClauseContext)

	// EnterReadingClause is called when entering the readingClause production.
	EnterReadingClause(c *ReadingClauseContext)

	// EnterMatch is called when entering the match production.
	EnterMatch(c *MatchContext)

	// EnterUnwind is called when entering the unwind production.
	EnterUnwind(c *UnwindContext)

	// EnterMerge is called when entering the merge production.
	EnterMerge(c *MergeContext)

	// EnterMergeAction is called when entering the mergeAction production.
	EnterMergeAction(c *MergeActionContext)

	// EnterCreate is called when entering the create production.
	EnterCreate(c *CreateContext)

	// EnterSet is called when entering the set production.
	EnterSet(c *SetContext)

	// EnterSetItem is called when entering the setItem production.
	EnterSetItem(c *SetItemContext)

	// EnterDelete is called when entering the delete production.
	EnterDelete(c *DeleteContext)

	// EnterRemove is called when entering the remove production.
	EnterRemove(c *RemoveContext)

	// EnterRemoveItem is called when entering the removeItem production.
	EnterRemoveItem(c *RemoveItemContext)

	// EnterInQueryCall is called when entering the inQueryCall production.
	EnterInQueryCall(c *InQueryCallContext)

	// EnterStandaloneCall is called when entering the standaloneCall production.
	EnterStandaloneCall(c *StandaloneCallContext)

	// EnterYieldItems is called when entering the yieldItems production.
	EnterYieldItems(c *YieldItemsContext)

	// EnterYieldItem is called when entering the yieldItem production.
	EnterYieldItem(c *YieldItemContext)

	// EnterWith is called when entering the with production.
	EnterWith(c *WithContext)

	// EnterReturn is called when entering the return production.
	EnterReturn(c *ReturnContext)

	// EnterReturnBody is called when entering the returnBody production.
	EnterReturnBody(c *ReturnBodyContext)

	// EnterReturnItems is called when entering the returnItems production.
	EnterReturnItems(c *ReturnItemsContext)

	// EnterReturnItem is called when entering the returnItem production.
	EnterReturnItem(c *ReturnItemContext)

	// EnterOrder is called when entering the order production.
	EnterOrder(c *OrderContext)

	// EnterSkip is called when entering the skip production.
	EnterSkip(c *SkipContext)

	// EnterLimit is called when entering the limit production.
	EnterLimit(c *LimitContext)

	// EnterSortItem is called when entering the sortItem production.
	EnterSortItem(c *SortItemContext)

	// EnterWhere is called when entering the where production.
	EnterWhere(c *WhereContext)

	// EnterPattern is called when entering the pattern production.
	EnterPattern(c *PatternContext)

	// EnterPatternPart is called when entering the patternPart production.
	EnterPatternPart(c *PatternPartContext)

	// EnterAnonymousPatternPart is called when entering the anonymousPatternPart production.
	EnterAnonymousPatternPart(c *AnonymousPatternPartContext)

	// EnterPatternElement is called when entering the patternElement production.
	EnterPatternElement(c *PatternElementContext)

	// EnterNodePattern is called when entering the nodePattern production.
	EnterNodePattern(c *NodePatternContext)

	// EnterPatternElementChain is called when entering the patternElementChain production.
	EnterPatternElementChain(c *PatternElementChainContext)

	// EnterRelationshipPattern is called when entering the relationshipPattern production.
	EnterRelationshipPattern(c *RelationshipPatternContext)

	// EnterRelationshipDetail is called when entering the relationshipDetail production.
	EnterRelationshipDetail(c *RelationshipDetailContext)

	// EnterProperties is called when entering the properties production.
	EnterProperties(c *PropertiesContext)

	// EnterRelationshipTypes is called when entering the relationshipTypes production.
	EnterRelationshipTypes(c *RelationshipTypesContext)

	// EnterNodeLabels is called when entering the nodeLabels production.
	EnterNodeLabels(c *NodeLabelsContext)

	// EnterNodeLabel is called when entering the nodeLabel production.
	EnterNodeLabel(c *NodeLabelContext)

	// EnterRangeLiteral is called when entering the rangeLiteral production.
	EnterRangeLiteral(c *RangeLiteralContext)

	// EnterLabelName is called when entering the labelName production.
	EnterLabelName(c *LabelNameContext)

	// EnterRelTypeName is called when entering the relTypeName production.
	EnterRelTypeName(c *RelTypeNameContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterOrExpression is called when entering the orExpression production.
	EnterOrExpression(c *OrExpressionContext)

	// EnterXorExpression is called when entering the xorExpression production.
	EnterXorExpression(c *XorExpressionContext)

	// EnterAndExpression is called when entering the andExpression production.
	EnterAndExpression(c *AndExpressionContext)

	// EnterNotExpression is called when entering the notExpression production.
	EnterNotExpression(c *NotExpressionContext)

	// EnterComparisonExpression is called when entering the comparisonExpression production.
	EnterComparisonExpression(c *ComparisonExpressionContext)

	// EnterAddOrSubtractExpression is called when entering the addOrSubtractExpression production.
	EnterAddOrSubtractExpression(c *AddOrSubtractExpressionContext)

	// EnterMultiplyDivideModuloExpression is called when entering the multiplyDivideModuloExpression production.
	EnterMultiplyDivideModuloExpression(c *MultiplyDivideModuloExpressionContext)

	// EnterPowerOfExpression is called when entering the powerOfExpression production.
	EnterPowerOfExpression(c *PowerOfExpressionContext)

	// EnterUnaryAddOrSubtractExpression is called when entering the unaryAddOrSubtractExpression production.
	EnterUnaryAddOrSubtractExpression(c *UnaryAddOrSubtractExpressionContext)

	// EnterStringListNullOperatorExpression is called when entering the stringListNullOperatorExpression production.
	EnterStringListNullOperatorExpression(c *StringListNullOperatorExpressionContext)

	// EnterPropertyOrLabelsExpression is called when entering the propertyOrLabelsExpression production.
	EnterPropertyOrLabelsExpression(c *PropertyOrLabelsExpressionContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterBooleanLiteral is called when entering the booleanLiteral production.
	EnterBooleanLiteral(c *BooleanLiteralContext)

	// EnterListLiteral is called when entering the listLiteral production.
	EnterListLiteral(c *ListLiteralContext)

	// EnterPartialComparisonExpression is called when entering the partialComparisonExpression production.
	EnterPartialComparisonExpression(c *PartialComparisonExpressionContext)

	// EnterParenthesizedExpression is called when entering the parenthesizedExpression production.
	EnterParenthesizedExpression(c *ParenthesizedExpressionContext)

	// EnterRelationshipsPattern is called when entering the relationshipsPattern production.
	EnterRelationshipsPattern(c *RelationshipsPatternContext)

	// EnterFilterExpression is called when entering the filterExpression production.
	EnterFilterExpression(c *FilterExpressionContext)

	// EnterIdInColl is called when entering the idInColl production.
	EnterIdInColl(c *IdInCollContext)

	// EnterFunctionInvocation is called when entering the functionInvocation production.
	EnterFunctionInvocation(c *FunctionInvocationContext)

	// EnterFunctionName is called when entering the functionName production.
	EnterFunctionName(c *FunctionNameContext)

	// EnterExplicitProcedureInvocation is called when entering the explicitProcedureInvocation production.
	EnterExplicitProcedureInvocation(c *ExplicitProcedureInvocationContext)

	// EnterImplicitProcedureInvocation is called when entering the implicitProcedureInvocation production.
	EnterImplicitProcedureInvocation(c *ImplicitProcedureInvocationContext)

	// EnterProcedureResultField is called when entering the procedureResultField production.
	EnterProcedureResultField(c *ProcedureResultFieldContext)

	// EnterProcedureName is called when entering the procedureName production.
	EnterProcedureName(c *ProcedureNameContext)

	// EnterNamespace is called when entering the namespace production.
	EnterNamespace(c *NamespaceContext)

	// EnterListComprehension is called when entering the listComprehension production.
	EnterListComprehension(c *ListComprehensionContext)

	// EnterPatternComprehension is called when entering the patternComprehension production.
	EnterPatternComprehension(c *PatternComprehensionContext)

	// EnterPropertyLookup is called when entering the propertyLookup production.
	EnterPropertyLookup(c *PropertyLookupContext)

	// EnterCaseExpression is called when entering the caseExpression production.
	EnterCaseExpression(c *CaseExpressionContext)

	// EnterCaseAlternatives is called when entering the caseAlternatives production.
	EnterCaseAlternatives(c *CaseAlternativesContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterNumberLiteral is called when entering the numberLiteral production.
	EnterNumberLiteral(c *NumberLiteralContext)

	// EnterMapLiteral is called when entering the mapLiteral production.
	EnterMapLiteral(c *MapLiteralContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterPropertyExpression is called when entering the propertyExpression production.
	EnterPropertyExpression(c *PropertyExpressionContext)

	// EnterPropertyKeyName is called when entering the propertyKeyName production.
	EnterPropertyKeyName(c *PropertyKeyNameContext)

	// EnterIntegerLiteral is called when entering the integerLiteral production.
	EnterIntegerLiteral(c *IntegerLiteralContext)

	// EnterDoubleLiteral is called when entering the doubleLiteral production.
	EnterDoubleLiteral(c *DoubleLiteralContext)

	// EnterSchemaName is called when entering the schemaName production.
	EnterSchemaName(c *SchemaNameContext)

	// EnterReservedWord is called when entering the reservedWord production.
	EnterReservedWord(c *ReservedWordContext)

	// EnterSymbolicName is called when entering the symbolicName production.
	EnterSymbolicName(c *SymbolicNameContext)

	// EnterLeftArrowHead is called when entering the leftArrowHead production.
	EnterLeftArrowHead(c *LeftArrowHeadContext)

	// EnterRightArrowHead is called when entering the rightArrowHead production.
	EnterRightArrowHead(c *RightArrowHeadContext)

	// EnterDash is called when entering the dash production.
	EnterDash(c *DashContext)

	// ExitCypher is called when exiting the cypher production.
	ExitCypher(c *CypherContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitRegularQuery is called when exiting the regularQuery production.
	ExitRegularQuery(c *RegularQueryContext)

	// ExitUnion is called when exiting the union production.
	ExitUnion(c *UnionContext)

	// ExitSingleQuery is called when exiting the singleQuery production.
	ExitSingleQuery(c *SingleQueryContext)

	// ExitSinglePartQuery is called when exiting the singlePartQuery production.
	ExitSinglePartQuery(c *SinglePartQueryContext)

	// ExitReadOnlyEnd is called when exiting the readOnlyEnd production.
	ExitReadOnlyEnd(c *ReadOnlyEndContext)

	// ExitReadUpdateEnd is called when exiting the readUpdateEnd production.
	ExitReadUpdateEnd(c *ReadUpdateEndContext)

	// ExitUpdatingEnd is called when exiting the updatingEnd production.
	ExitUpdatingEnd(c *UpdatingEndContext)

	// ExitMultiPartQuery is called when exiting the multiPartQuery production.
	ExitMultiPartQuery(c *MultiPartQueryContext)

	// ExitReadPart is called when exiting the readPart production.
	ExitReadPart(c *ReadPartContext)

	// ExitUpdatingPart is called when exiting the updatingPart production.
	ExitUpdatingPart(c *UpdatingPartContext)

	// ExitUpdatingStartClause is called when exiting the updatingStartClause production.
	ExitUpdatingStartClause(c *UpdatingStartClauseContext)

	// ExitUpdatingClause is called when exiting the updatingClause production.
	ExitUpdatingClause(c *UpdatingClauseContext)

	// ExitReadingClause is called when exiting the readingClause production.
	ExitReadingClause(c *ReadingClauseContext)

	// ExitMatch is called when exiting the match production.
	ExitMatch(c *MatchContext)

	// ExitUnwind is called when exiting the unwind production.
	ExitUnwind(c *UnwindContext)

	// ExitMerge is called when exiting the merge production.
	ExitMerge(c *MergeContext)

	// ExitMergeAction is called when exiting the mergeAction production.
	ExitMergeAction(c *MergeActionContext)

	// ExitCreate is called when exiting the create production.
	ExitCreate(c *CreateContext)

	// ExitSet is called when exiting the set production.
	ExitSet(c *SetContext)

	// ExitSetItem is called when exiting the setItem production.
	ExitSetItem(c *SetItemContext)

	// ExitDelete is called when exiting the delete production.
	ExitDelete(c *DeleteContext)

	// ExitRemove is called when exiting the remove production.
	ExitRemove(c *RemoveContext)

	// ExitRemoveItem is called when exiting the removeItem production.
	ExitRemoveItem(c *RemoveItemContext)

	// ExitInQueryCall is called when exiting the inQueryCall production.
	ExitInQueryCall(c *InQueryCallContext)

	// ExitStandaloneCall is called when exiting the standaloneCall production.
	ExitStandaloneCall(c *StandaloneCallContext)

	// ExitYieldItems is called when exiting the yieldItems production.
	ExitYieldItems(c *YieldItemsContext)

	// ExitYieldItem is called when exiting the yieldItem production.
	ExitYieldItem(c *YieldItemContext)

	// ExitWith is called when exiting the with production.
	ExitWith(c *WithContext)

	// ExitReturn is called when exiting the return production.
	ExitReturn(c *ReturnContext)

	// ExitReturnBody is called when exiting the returnBody production.
	ExitReturnBody(c *ReturnBodyContext)

	// ExitReturnItems is called when exiting the returnItems production.
	ExitReturnItems(c *ReturnItemsContext)

	// ExitReturnItem is called when exiting the returnItem production.
	ExitReturnItem(c *ReturnItemContext)

	// ExitOrder is called when exiting the order production.
	ExitOrder(c *OrderContext)

	// ExitSkip is called when exiting the skip production.
	ExitSkip(c *SkipContext)

	// ExitLimit is called when exiting the limit production.
	ExitLimit(c *LimitContext)

	// ExitSortItem is called when exiting the sortItem production.
	ExitSortItem(c *SortItemContext)

	// ExitWhere is called when exiting the where production.
	ExitWhere(c *WhereContext)

	// ExitPattern is called when exiting the pattern production.
	ExitPattern(c *PatternContext)

	// ExitPatternPart is called when exiting the patternPart production.
	ExitPatternPart(c *PatternPartContext)

	// ExitAnonymousPatternPart is called when exiting the anonymousPatternPart production.
	ExitAnonymousPatternPart(c *AnonymousPatternPartContext)

	// ExitPatternElement is called when exiting the patternElement production.
	ExitPatternElement(c *PatternElementContext)

	// ExitNodePattern is called when exiting the nodePattern production.
	ExitNodePattern(c *NodePatternContext)

	// ExitPatternElementChain is called when exiting the patternElementChain production.
	ExitPatternElementChain(c *PatternElementChainContext)

	// ExitRelationshipPattern is called when exiting the relationshipPattern production.
	ExitRelationshipPattern(c *RelationshipPatternContext)

	// ExitRelationshipDetail is called when exiting the relationshipDetail production.
	ExitRelationshipDetail(c *RelationshipDetailContext)

	// ExitProperties is called when exiting the properties production.
	ExitProperties(c *PropertiesContext)

	// ExitRelationshipTypes is called when exiting the relationshipTypes production.
	ExitRelationshipTypes(c *RelationshipTypesContext)

	// ExitNodeLabels is called when exiting the nodeLabels production.
	ExitNodeLabels(c *NodeLabelsContext)

	// ExitNodeLabel is called when exiting the nodeLabel production.
	ExitNodeLabel(c *NodeLabelContext)

	// ExitRangeLiteral is called when exiting the rangeLiteral production.
	ExitRangeLiteral(c *RangeLiteralContext)

	// ExitLabelName is called when exiting the labelName production.
	ExitLabelName(c *LabelNameContext)

	// ExitRelTypeName is called when exiting the relTypeName production.
	ExitRelTypeName(c *RelTypeNameContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitOrExpression is called when exiting the orExpression production.
	ExitOrExpression(c *OrExpressionContext)

	// ExitXorExpression is called when exiting the xorExpression production.
	ExitXorExpression(c *XorExpressionContext)

	// ExitAndExpression is called when exiting the andExpression production.
	ExitAndExpression(c *AndExpressionContext)

	// ExitNotExpression is called when exiting the notExpression production.
	ExitNotExpression(c *NotExpressionContext)

	// ExitComparisonExpression is called when exiting the comparisonExpression production.
	ExitComparisonExpression(c *ComparisonExpressionContext)

	// ExitAddOrSubtractExpression is called when exiting the addOrSubtractExpression production.
	ExitAddOrSubtractExpression(c *AddOrSubtractExpressionContext)

	// ExitMultiplyDivideModuloExpression is called when exiting the multiplyDivideModuloExpression production.
	ExitMultiplyDivideModuloExpression(c *MultiplyDivideModuloExpressionContext)

	// ExitPowerOfExpression is called when exiting the powerOfExpression production.
	ExitPowerOfExpression(c *PowerOfExpressionContext)

	// ExitUnaryAddOrSubtractExpression is called when exiting the unaryAddOrSubtractExpression production.
	ExitUnaryAddOrSubtractExpression(c *UnaryAddOrSubtractExpressionContext)

	// ExitStringListNullOperatorExpression is called when exiting the stringListNullOperatorExpression production.
	ExitStringListNullOperatorExpression(c *StringListNullOperatorExpressionContext)

	// ExitPropertyOrLabelsExpression is called when exiting the propertyOrLabelsExpression production.
	ExitPropertyOrLabelsExpression(c *PropertyOrLabelsExpressionContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitBooleanLiteral is called when exiting the booleanLiteral production.
	ExitBooleanLiteral(c *BooleanLiteralContext)

	// ExitListLiteral is called when exiting the listLiteral production.
	ExitListLiteral(c *ListLiteralContext)

	// ExitPartialComparisonExpression is called when exiting the partialComparisonExpression production.
	ExitPartialComparisonExpression(c *PartialComparisonExpressionContext)

	// ExitParenthesizedExpression is called when exiting the parenthesizedExpression production.
	ExitParenthesizedExpression(c *ParenthesizedExpressionContext)

	// ExitRelationshipsPattern is called when exiting the relationshipsPattern production.
	ExitRelationshipsPattern(c *RelationshipsPatternContext)

	// ExitFilterExpression is called when exiting the filterExpression production.
	ExitFilterExpression(c *FilterExpressionContext)

	// ExitIdInColl is called when exiting the idInColl production.
	ExitIdInColl(c *IdInCollContext)

	// ExitFunctionInvocation is called when exiting the functionInvocation production.
	ExitFunctionInvocation(c *FunctionInvocationContext)

	// ExitFunctionName is called when exiting the functionName production.
	ExitFunctionName(c *FunctionNameContext)

	// ExitExplicitProcedureInvocation is called when exiting the explicitProcedureInvocation production.
	ExitExplicitProcedureInvocation(c *ExplicitProcedureInvocationContext)

	// ExitImplicitProcedureInvocation is called when exiting the implicitProcedureInvocation production.
	ExitImplicitProcedureInvocation(c *ImplicitProcedureInvocationContext)

	// ExitProcedureResultField is called when exiting the procedureResultField production.
	ExitProcedureResultField(c *ProcedureResultFieldContext)

	// ExitProcedureName is called when exiting the procedureName production.
	ExitProcedureName(c *ProcedureNameContext)

	// ExitNamespace is called when exiting the namespace production.
	ExitNamespace(c *NamespaceContext)

	// ExitListComprehension is called when exiting the listComprehension production.
	ExitListComprehension(c *ListComprehensionContext)

	// ExitPatternComprehension is called when exiting the patternComprehension production.
	ExitPatternComprehension(c *PatternComprehensionContext)

	// ExitPropertyLookup is called when exiting the propertyLookup production.
	ExitPropertyLookup(c *PropertyLookupContext)

	// ExitCaseExpression is called when exiting the caseExpression production.
	ExitCaseExpression(c *CaseExpressionContext)

	// ExitCaseAlternatives is called when exiting the caseAlternatives production.
	ExitCaseAlternatives(c *CaseAlternativesContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitNumberLiteral is called when exiting the numberLiteral production.
	ExitNumberLiteral(c *NumberLiteralContext)

	// ExitMapLiteral is called when exiting the mapLiteral production.
	ExitMapLiteral(c *MapLiteralContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitPropertyExpression is called when exiting the propertyExpression production.
	ExitPropertyExpression(c *PropertyExpressionContext)

	// ExitPropertyKeyName is called when exiting the propertyKeyName production.
	ExitPropertyKeyName(c *PropertyKeyNameContext)

	// ExitIntegerLiteral is called when exiting the integerLiteral production.
	ExitIntegerLiteral(c *IntegerLiteralContext)

	// ExitDoubleLiteral is called when exiting the doubleLiteral production.
	ExitDoubleLiteral(c *DoubleLiteralContext)

	// ExitSchemaName is called when exiting the schemaName production.
	ExitSchemaName(c *SchemaNameContext)

	// ExitReservedWord is called when exiting the reservedWord production.
	ExitReservedWord(c *ReservedWordContext)

	// ExitSymbolicName is called when exiting the symbolicName production.
	ExitSymbolicName(c *SymbolicNameContext)

	// ExitLeftArrowHead is called when exiting the leftArrowHead production.
	ExitLeftArrowHead(c *LeftArrowHeadContext)

	// ExitRightArrowHead is called when exiting the rightArrowHead production.
	ExitRightArrowHead(c *RightArrowHeadContext)

	// ExitDash is called when exiting the dash production.
	ExitDash(c *DashContext)
}
