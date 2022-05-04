package types

import (
	"testing"

	"github.com/Karan-3108/Fortress/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCancelFortress_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCancelFortress
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCancelFortress{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCancelFortress{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
