package optimizer

// PlayerStatRow TODO: write me
// order is VERY important
type PlayerStatRow struct {
	ID                             int     // A  - " "
	PlayerName                     string  // B  - "Player"
	TeamAbbrv                      string  // C  - "Tm"
	Age                            int     // D  - "Age"
	Position                       string  // E  - "Pos"
	GamesPlayed                    int     // F  - "G"
	GamesStarted                   int     // G  - "GS"
	PassCmpl                       int     // H  - "pass_Cmp"
	PassAtt                        int     // I  - "pass_Att"
	PassCmplPct                    float32 // J  - "pass_Cmp%"
	PassYards                      int     // K  - "pass_Yds"
	PassTD                         int     // L  - "pass_TD"
	PassTDPct                      float32 // M  - "pass_TD%"
	PassInt                        int     // N  - "pass_Int"
	PassIntPct                     float32 // O  - "pass_Int%"
	PassYPA                        float32 // P  - "pass_Y/A"
	PassYPC                        float32 // Q  - "pass_Y/C"
	PassYPG                        float32 // R  - "pass_Y/G"
	PassSacks                      int     // S  - "pass_Sk"
	PassSackPct                    float32 // T  - "pass_Sk%"
	RushAtt                        int     // U  - "rush_Att"
	RushYards                      int     // V  - "rush_Yds"
	RushTD                         int     // W  - "rush_TD"
	RushYPA                        float32 // X  - "rush_Y/A"
	RushYPG                        float32 // Y  - "rush_Y/G"
	RushAPG                        float32 // Z  - "rush_A/G"
	RushFum                        int     // AA - "rush_Fmb"
	Targets                        int     // AB - "rec_Tgt"
	Receptions                     int     // AC - "rec_Rec"
	CatchPct                       float32 // AD - "rec_Ctch%"
	RecYards                       int     // AE - "rec_Yds"
	RecYPR                         float32 // AF - "rec_Y/R" yards per reception
	RecTD                          int     // AG - "rec_TD"
	RecRPG                         float32 // AH - "rec_R/G" receptions per game
	RecYPG                         float32 // AI - "rec_Y/G"
	RecFum                         int     // AJ - "rec_Fmb"
	FGAtt0_19                      int     // AK - "kick_FGA" 0-19
	FGMde0_19                      int     // AL - "kick_FGM" 0-19
	FGAtt20_29                     int     // AM - "kick_FGA" 20-29
	FGMde20_29                     int     // AN - "kick_FGM" 20-29
	FGAtt30_39                     int     // AO - "kick_FGA" 30-39
	FGMde30_39                     int     // AP - "kick_FGM" 30-39
	FGAtt40_49                     int     // AQ - "kick_FGA" 40-49
	FGMde40_49                     int     // AR - "kick_FGM" 40-49
	FGAtt50Plus                    int     // AS - "kick_FGA" 50+
	FGMde50Plus                    int     // AT - "kick_FGM" 50+
	FGAttTotal                     int     // AU - "kick_FGA" Total
	FGMdeTotal                     int     // AV - "kick_FGM" Total
	FGTotalPct                     float32 // AW - "kick_FG%"
	XPAtt                          int     // AX - "kick_XPA"
	XPMade                         int     // AY - "kick_XPM"
	XPPct                          float32 // AZ - "kick_XP%"
	StandardFantasyPts             float32 // BA - "fan_FantPt" decimal?
	PPRFantasyPts                  float32 // BB - "fan_PPR" decimal?
	DraftKingsFPts                 float32 // BC - "fan_DKPt" decimal?
	FanDuelFPts                    float32 // BD - "fan_FDPt" decimal?
	ValueBasedDrafting             int     // BE - "fan_VBD" Value based drafting
	PositionalRank                 int     // BF - "fan_PosRank" positional rank? based upon what?
	OverallRank                    int     // BG - "fan_OvRank" overall rank? based upon what?
	RZPassCmplIn20                 int     // BH - "rz_pass_Cmp" Inside the 20
	RZPassAttIn20                  int     // BI - "rz_pass_Att" Inside the 20
	RZPassCmplPctIn20              float32 // BJ - "rz_pass_Cmp%" Inside the 20
	RZPassYardsIn20                int     // BK - "rz_pass_Yds" Inside the 20
	RZPassTDIn20                   int     // BL - "rz_pass_TD" Inside the 20
	RZPassIntIn20                  int     // BM - "rz_pass_Int" Inside the 20
	RZPassCmplIn10                 int     // BN - "rz_pass_Cmp" Inside the 10
	RZPassAttIn10                  int     // BO - "rz_pass_Att" Inside the 10
	RZPassCmplPctIn10              float32 // BP - "rz_pass_Cmp%" Inside the 10
	RZPassYardsIn10                int     // BQ - "rz_pass_Yds" Inside the 10
	RZPassTDIn10                   int     // BR - "rz_pass_TD" Inside the 10
	RZPassIntIn10                  int     // BS - "rz_pass_Int" Inside the 10
	RZRushAttIn20                  int     // BT - "rz_rush_Att" Inside the 20
	RZRushYardsIn20                int     // BU - "rz_rush_Yds" Inside the 20
	RZRushTDIn20                   int     // BV - "rz_rush_TD" Inside the 20
	RZRushShareIn20                float32 // BW - "rz_rush_%Rush" Inside the 20 TODO: maybe a string?
	RZRushAttIn10                  int     // BX - "rz_rush_Att" Inside the 10
	RZRushYardsIn10                int     // BY - "rz_rush_Yds" Inside the 10
	RZRushTDIn10                   int     // BZ - "rz_rush_TD" Inside the 10
	RZRushShareIn10                float32 // CA - "rz_rush_%Rush" Inside the 10 TODO: maybe a string?
	RZRushAttIn5                   int     // CB - "rz_rush_Att" Inside the 5
	RZRushYardsIn5                 int     // CC - "rz_rush_Yds" Inside the 5
	RZRushTDIn5                    int     // CD - "rz_rush_TD" Inside the 5
	RZRushShareIn5                 float32 // CE - "rz_rush_%Rush" Inside the 5 TODO: maybe a string?
	RZTargetsIn20                  int     // CF - "rz_rec_Tgt"
	RZReceptionsIn20               int     // CG - "rz_rec_Rec"
	RZCatchPctIn20                 float32 // CH - "rz_rec_Ctch%"
	RZRecYardsIn20                 int     // CI - "rz_rec_Yds"
	RZRecTDIn20                    int     // CJ - "rz_rec_TD"
	RZTargetShareIn20              float32 // CK - "rz_rec_%Tgt"
	RZTargetsIn10                  int     // CL - "rz_rec_Tgt"
	RZReceptionsIn10               int     // CM - "rz_rec_Rec"
	RZCatchPctIn10                 float32 // CN - "rz_rec_Ctch%"
	RZRecYardsIn10                 int     // CO - "rz_rec_Yds"
	RZRecTDIn10                    int     // CP - "rz_rec_TD"
	RZTargetShareIn10              float32 // CQ - "rz_rec_%Tgt"
	TeamDefPtsAllowed              int     // CR - "team_d_PF"
	TeamDefYardsAllowed            int     // CS - "team_d_Yds"
	TeamDefPlays                   int     // CT - "team_d_Ply"
	TeamDefYPP                     float32 // CU - "team_d_Y/P" yards per play
	TeamDefTurnovers               int     // CV - "team_d_TO"
	TeamDefFumblesRec              int     // CW - "team_d_FL"
	TeamDefFirstDowns              int     // CX - "team_d_1stD"
	TeamDefCmplAllowed             int     // CY - "team_d_Cmp"
	TeamDefPassingAttempts         int     // CZ - "team_d_Att"  ??????
	TeamDefPassingYards            int     // DA - "team_d_Yds"  ??????
	TeamDefPassingTD               int     // DB - "team_d_TD"   ??????
	TeamDefInt                     int     // DC - "team_d_Int"  ??????
	TeamDefNetPassingYPA           float32 // DD - "team_d_NY/A" ??????
	TeamDefPassingFirstDowns       int     // DE - "team_d_1stD" ??????
	TeamDefRushingAttempts         int     // DF - "team_d_Att"  ??????
	TeamDefRushingYards            int     // DG - "team_d_Yds"  ??????
	TeamDefRushingTD               int     // DH - "team_d_TD"   ??????
	TeamDefRushingYPA              float32 // DI - "team_d_Y/A" ???????
	TeamDefRushingFirstDowns       int     // DJ - "team_d_1stD" ??????
	TeamDefPenalties               int     // DK - "team_d_Pen"
	TeamDefPenaltyYards            int     // DL - "team_d_Yds" ??????
	TeamDefFirstPY                 int     // DM - "team_d_1stPy" ?????? first downs given by penalty yards???
	TeamDefScoringPct              float32 // DN - "team_d_Sc%" ??????? scoring percentage? like how many drivens result in scores / total drives
	TeamDefTOPct                   float32 // DO - "team_d_TO%" ??????
	TeamDefEXP                     float32 // DP - "team_d_EXP" ????? what the hell is this??
	DefRunningBackAttempts         int     // DQ - "DvsRB_Att" ??????
	DefRunningBackYards            int     // DR - "DvsRB_Yds" ?????? rushing?
	DefRunningBackTD               int     // DS - "DvsRB_TD"  ?????? rushing?
	DefRunningBackTargets          int     // DT - "DvsRB_Tgt" ??????
	DefRunningBackReceptions       int     // DU - "DvsRB_Rec" ??????
	DefRunningBackYards2           int     // DV - "DvsRB_Yds" ?????? receiving?
	DefRunningBackTD2              int     // DW - "DvsRB_TD"  ?????? receiving?
	DefRunningBackPoints           float32 // DX - "DvsRB_FantPt" ????? fantasy points? what scoring type? standard maybe?
	DefRunningBackDraftKingsPoints float32 // DY - "DvsRB_DKPt" ????? draft kings points to running backs?
	DefRunningBackFanDuelPoints    float32 // DZ - "DvsRB_FDPt" ?????? fanduel fantasy points to running backs?
	DefVsTETargets                 int     // EA - "DvsTE_Tgt"
	DefVsTEReceptions              int     // EB - "DvsTE_Rec"
	DefVsTEYards                   int     // EC - "DvsTE_Yds"
	DefVsTETD                      int     // ED - "DvsTE_TD"
	DefVsTEFantasyPoints           float32 // EE - "DvsTE_FantPt" ???? what scoring??
	DefVsTEDraftKingsPoints        float32 // EF - "DvsTE_DKPt"
	DefVsTEFanDuelPoints           float32 // EG - "DvsTE_FDPt"
	DefVsQBCompletions             int     // EH - "DvsQB_Cmp"
	DefVsQBPassingAttempts         int     // EI - "DvsQB_Att"
	DefVsQBPassingYards            int     // EJ - "DvsQB_Yds"
	DefVsQBPassingTD               int     // EK - "DvsQB_TD"
	DefVsQBInterceptions           int     // EL - "DvsQB_Int"
	DefVsQB2PtConversion           int     // EM - "DvsQB_2PP" ????????? is this a pass or rush??
	DefVsQBSacks                   int     // EN - "DvsQB_Sk"
	DefVsQBRushingAttempts         int     // EO - "DvsQB_Att"
	DefVsQBRushingYards            int     // EP - "DvsQB_Yds"
	DefVsQBRushingTD               int     // EQ - "DvsQB_TD"
	DefVsQBFantasyPoints           float32 // ER - "DvsQB_FantPt" ????? what scoring??
	DefVsQBDraftKingsPoints        float32 // ES - "DvsQB_DKPt"
	DefVsQBFanDuelPoints           float32 // ET - "DvsQB_FDPt"
	DefVsWRTargets                 int     // EU - "DvsWR_Tgt"
	DefVsWRReceptions              int     // EV - "DvsWR_Rec"
	DefVsWRYards                   int     // EW - "DvsWR_Yds"
	DefVsWRTD                      int     // EX - "DvsWR_TD"
	DefVsWRFantasyPoints           int     // EY - "DvsWR_FantPt" ????? what scoring?
	DefVsWRDraftKingsPoints        int     // EZ - "DvsWR_DKPt"
	DefVsWRFanDuelPoints           int     // FA - "DvsWR_FDPt"
}
